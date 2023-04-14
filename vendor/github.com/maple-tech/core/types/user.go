package types

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"encoding/json"

	// "encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/maple-tech/core/db"
	"github.com/maple-tech/core/log"
)

// User represents the bare-essentials of a user's record in the system.
type User struct {
	ID      URID    `json:"id" db:"u_id"`
	Name    string  `json:"name" db:"u_name"`
	Picture *string `json:"picture" db:"u_picture"`
}

// UserRecord expands on user by adding the active, and time properties
type UserRecord struct {
	User
	Active      bool     `json:"active" db:"u_active"`
	TimeCreated NullTime `json:"timeCreated" db:"u_created"`
	TimeUpdated NullTime `json:"timeUpdated,omitempty" db:"u_updated"`
	TimeDeleted NullTime `json:"timeDeleted,omitempty" db:"u_deleted"`
}

// UserSettings holds the users global account settings
type UserSettings struct {
	Language      string   `json:"language" db:"us_lang"`
	TwoFactorMode int32    `json:"twoFactorMode" db:"us_tfm"`
	TimeUpdated   NullTime `json:"timeUpdated,omitempty" db:"us_updated"`
	// TimeUpdated   NullTime `json:"timeUpdated,omitempty" db:"us_updated"`
}

// UserDepartment links a user to a department subscription, including the department information
type UserDepartment struct {
	Department
	Supervisor bool `json:"supervisor" db:"ud_supr"`
	// TimeJoined time.Time `json:"-" db:"ud_created"`
	TimeJoined NullTime `json:"timeJoined" db:"ud_created"`
}

// UserPosition links a user to a position subscription, including the position information
type UserPosition struct {
	Position
	TimeJoined NullTime `json:"timeJoined" db:"up_created"`
	// TimeJoined time.Time `json:"-" db:"up_created"`
}
type CompanyModuleArray []CompanyModule

func (pc *CompanyModuleArray) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &pc)
		return nil
	case string:
		json.Unmarshal([]byte(v), &pc)
		return nil
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
func (pc *CompanyModuleArray) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

// UserCompany links a user to a membership of a company
type UserCompany struct {
	CompanyRecord
	Modules          []CompanyModule `json:"modules" db:"modules"`
	MembershipActive bool            `json:"membershipActive" db:"uc_active"`
	TimeJoined       NullTime        `json:"-" db:"uc_joined"`
	// TimeMembershipUpdated NullTime        `json:"-" db:"uc_updated"`
	TimeMembershipUpdated  NullTime         `json:"timeMembershipUpdated,omitempty" db:"uc_updated"`
	Role                   PermissionModal  `json:"role" db:"role"`
	RoleID                 ID               `json:"-" db:"uc_role"`
	Departments            []UserDepartment `json:"departments" db:"departments"`
	Positions              []UserPosition   `json:"positions" db:"positions"`
	IsExternal             bool             `json:"isExternal" db:"uc_external"`
	ExternalCompany        *Company         `json:"externalCompany"`
	ExternalCompanyID      *URID            `json:"-" db:"uce_id"`
	ExternalCompanyName    *string          `json:"-" db:"uce_name"`
	ExternalCompanyPicture *string          `json:"-" db:"uce_picture"`
}

// CheckPermissions takes a variable amount of permissions and checks that the user
// has all the permissions needed. Returns a bool if all pass, and a slice of the ones missing.
func (u UserCompany) CheckPermissions(perms ...Permission) (bool, []Permission) {
	//Allow owners and super_admin users to pass through without full check
	if u.Role.Permissions.Has(PermOwner) || u.Role.Permissions.Has(PermSuperAdmin) {
		return true, []Permission{}
	}

	//Check permissions for NON-admin users
	missing := make([]Permission, 0)
	for i := range perms {
		if !u.Role.Permissions.Has(perms[i]) {
			missing = append(missing, perms[i])
		}
	}

	return len(missing) < 1, missing
}

// GetModule finds and returns a module within the company.
// Returns nil if not found
func (u UserCompany) GetModule(key string) *CompanyModule {
	for _, mod := range u.Modules {
		if mod.Module == key {
			return &mod
		}
	}
	return nil
}

type UserJSONSettings UserSettings
type UserCompaniesList []UserCompany

func (pc *UserJSONSettings) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &pc)
	case string:
		return json.Unmarshal([]byte(v), &pc)
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
func (pc *UserJSONSettings) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

func (pc *UserCompaniesList) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		fmt.Println(string(v))
		return json.Unmarshal(v, &pc)
	case string:
		fmt.Println(string(v))

		return json.Unmarshal([]byte(v), &pc)

	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
func (pc *UserCompaniesList) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

// UserComplete is the full(ish) record for the user
type UserComplete struct {
	UserRecord
	Settings UserJSONSettings `json:"settings" db:"settings"`
	// Settings  UserSettings  `json:"settings"`
	Companies UserCompaniesList `json:"companies" db:"companies"`
	// Companies []UserCompany `json:"companies"`
}

// GetCompany finds and returns the company by provided ID within this users memberships.
// If no memberships, "no company memberships" error is returned.
// If no matches, "no matching company" error is returned.
func (u UserComplete) GetCompany(id URID) (*UserCompany, error) {
	if len(u.Companies) == 0 {
		return nil, errors.New("no company memberships")
	}

	for i := range u.Companies {
		if u.Companies[i].ID == id {
			return &(u.Companies[i]), nil
		}
	}

	return nil, errors.New("no matching company")
}

// GetUserCompleteByID returns the complete user object (UserComplete) based on the provided ID.
// Does this under a transaction for batching over the db connection.
//
// DEVELOPMENT NOTE: This currently does at minimum 3 calls, and 4 per company. Surely we can find a better way
func GetUserCompleteByID(id URID) (UserComplete, error) {
	var user UserComplete

	err := db.WrapTX(func(tx *sqlx.Tx) (e error) {
		//Grab the user object
		// TODO - This is a bit of a hack, but it's the only way to get the user record without the password
		return tx.Get(&user, `
SELECT id                                 AS u_id,
       name                               AS u_name,
       picture                            AS u_picture,
       active                             AS u_active,
       time_created                       AS u_created,
       time_updated                       AS u_updated,
       time_deleted                       AS u_deleted,
       TO_JSONB(settings.settings)        AS settings,
       ARRAY_TO_JSON(companies_bulk.list) AS companies
FROM core.users
	     LEFT JOIN LATERAL (
	SELECT JSON_BUILD_OBJECT(
			       'language', cus.language,
			       'twoFactorMode', cus.twofactor,
			       'timeUpdated', cus.time_updated::timestamptz
		       ) AS settings
	FROM core.user_settings cus
	WHERE cus.user_id = users.id
	) settings ON TRUE
	     LEFT JOIN LATERAL (
	SELECT ARRAY(
			       SELECT JSON_BUILD_OBJECT(
					              'id', c.id,
					              'infoId', ci.id,
					              'name', c.name,
					              'uniqueName', c.unique_name,
					              'picture', c.picture,
					              'industries', ci.industries,
					              'employees', ci.employees,
					              'contactInfo', ci.contact_info,
					              'membershipActive', uc.active,
-- 					              'uc_role', uc.permission_id,
					              'timeJoined', uc.time_created::timestamptz,
					              'timeMembershipUpdated', uc.time_updated::timestamptz,
				              -- 					              'ci_updated', ci.time_updated,
-- 					              'deleted', c.time_deleted,
					              'role', TO_JSON(p.*),
					              'modules', ARRAY_TO_JSON(modules.list),
					              'positions', ARRAY_TO_JSON(positions_list.list),
					              'departments', ARRAY_TO_JSON(departments.list)
				              )
			       FROM core.user_companies uc
				            JOIN core.companies c ON c.id = uc.company_id
				            JOIN core.company_infos ci ON ci.company_id = uc.company_id
				            JOIN core.company_permissions p ON p.id = uc.permission_id AND p.company_id = uc.company_id
				            LEFT JOIN LATERAL (SELECT ARRAY(
						                                      SELECT JSONB_BUILD_OBJECT(
								                                             'id', cm.id,
								                                             'module', cm.module,
								                                             'settings', cm.settings,
								                                             'active', cm.active,
								                                             'endOfPeriod', cm.end_of_period::timestamptz,
								                                             'timeCreated', cm.time_created::timestamptz,
								                                             'timeUpdate', cm.time_updated::timestamptz
							                                             )
						                                      FROM core.company_modules cm
						                                      WHERE cm.company_id = uc.company_id
					                                      ) list ) modules ON TRUE
				            LEFT JOIN LATERAL (
				       SELECT ARRAY(
						              SELECT JSON_BUILD_OBJECT(
								                     'id', p.id,
								                     'name', p.name,
								                     'uniqueName', p.unique_name,
								                     'timeCreated', p.time_created::timestamptz,
								                     'timeJoined', up.time_created::timestamptz
							                     )
						              FROM core.user_positions up
							                   JOIN core.positions p
							                        ON p.id = up.position_id AND p.company_id = up.company_id
						              WHERE up.user_id = uc.user_id
							            AND up.company_id = uc.company_id
					              ) list
				       ) AS positions_list ON TRUE
				            LEFT JOIN LATERAL (
				       SELECT ARRAY(
						              SELECT JSONB_BUILD_OBJECT(
								                     'id', d.id,
								                     'name', d.name,
								                     'uniqueName', d.unique_name,
								                     'timeCreated', d.time_created::timestamptz,
								                     'supervisor', ud.supervisor,
								                     'timeJoined', ud.time_created::timestamptz
							                     )
						              FROM core.user_departments ud
							                   JOIN core.departments d
							                        ON d.id = ud.department_id AND d.company_id = ud.company_id
						              WHERE ud.user_id = uc.user_id
							            AND ud.company_id = uc.company_id
					              ) AS list
				       ) AS departments ON TRUE
			       WHERE uc.user_id = users.id
		       ) list
	) AS companies_bulk ON TRUE
WHERE id =$1
		`, id)

	})

	return user, err
}
func GetUserCompleteByIDSlowly(id URID) (UserComplete, error) {
	var user UserComplete
	err := db.WrapTX(func(tx *sqlx.Tx) (e error) {
		//Grab the user object
		e = tx.Get(&user, `SELECT id as u_id, name as u_name, picture as u_picture,
			active as u_active, time_created as u_created, time_updated as u_updated, time_deleted as u_deleted
			FROM core.users
			WHERE id=$1`, id)
		if e != nil {
			log.Err("failed to get user object", e)
			return fmt.Errorf("failed to get user object from db; error = %s", e.Error())
		}

		//Fill in the settings for the user
		//NOTe: I would merge this into the above query to reduce 1 call
		e = tx.Get(&user.Settings, `SELECT language as us_lang, twofactor as us_tfm, time_updated as us_updated
			FROM core.user_settings
			WHERE user_id=$1`, user.ID)
		if e != nil {
			return fmt.Errorf("failed to get user settings from db; error = %s", e.Error())
		}

		//Grab the user company listings
		oe := tx.Select(&user.Companies, `SELECT uc.active as uc_active, uc.permission_id as uc_role, uc.time_created as uc_joined, uc.time_updated as uc_updated,
			c.id as c_id, c.name as c_name, c.unique_name as c_unique, c.picture as c_picture,
			ci.id as ci_id, ci.employees as ci_emp, ci.industries as ci_ind, ci.contact_info as ci_contact, ci.time_updated as ci_updated
			FROM core.user_companies uc
			JOIN core.companies c ON c.id=uc.company_id
			JOIN core.company_infos ci ON ci.company_id=uc.company_id
			WHERE uc.user_id=$1`, user.ID)
		//oe is an "optional" error since it might just be empty
		if oe != nil && oe != sql.ErrNoRows {
			return fmt.Errorf("failed to get user companies from db; error = %s", oe.Error())
		}

		//Fill in the missing data like modules, departments, position, and permission role, FOR EACH company
		//NOTE: To any DB pro's, any way to speed more of this up, please do
		for i := range user.Companies {
			//Grab the permission model for this company entry
			//NOTE: I would merge this into the query for getting the company to reduce 1 call
			if user.Companies[i].RoleID.Valid() {
				e = tx.Get(&(user.Companies[i].Role), `SELECT p.id as perm_id, p.name as perm_name, p.permissions as perm_perms,
					p.time_created as perm_created, p.time_updated as perm_updated
					FROM core.company_permissions p
					WHERE p.id=$1 AND p.company_id=$2`, user.Companies[i].RoleID, user.Companies[i].ID)
				if e != nil {
					return fmt.Errorf("failed to get permission for company (index %d); error = %s", i, e.Error())
				}
			}

			//Grab the modules for this company
			oe = tx.Select(&(user.Companies[i].Modules), `SELECT id as cm_id, module as cm_mod, settings as cm_settings, active as cm_active,
				end_of_period as cm_eop, time_created as cm_created, time_updated as cm_updated
				FROM core.company_modules
				WHERE company_id=$1`, user.Companies[i].ID)
			if oe != nil && oe != sql.ErrNoRows {
				return fmt.Errorf("failed to get company modules from db; error = %s", oe.Error())
			}

			//Fill in the user positions within this company
			oe = tx.Select(&(user.Companies[i].Positions), `SELECT p.id as p_id, p.name as p_name, p.unique_name as p_unique, p.time_created as p_created,
				up.time_created as up_created
				FROM core.user_positions up
				JOIN core.positions p ON p.id=up.position_id AND p.company_id=up.company_id
				WHERE up.user_id=$1 AND up.company_id=$2`, user.ID, user.Companies[i].ID)
			if oe != nil && oe != sql.ErrNoRows {
				return fmt.Errorf("failed to get positions for user within company from db; error = %s", oe.Error())
			}

			//Fill in the user departments within this company
			oe = tx.Select(&(user.Companies[i].Departments), `SELECT d.id as d_id, d.name as d_name, d.Unique_name as d_unique, d.time_created as d_created,
				ud.supervisor as ud_supr, ud.time_created as ud_created
				FROM core.user_departments ud
				JOIN core.departments d ON d.id=ud.department_id AND d.company_id=ud.company_id
				WHERE ud.user_id=$1 AND ud.company_id=$2`, user.ID, user.Companies[i].ID)
			if oe != nil && oe != sql.ErrNoRows {
				return fmt.Errorf("failed to get departments for user within company from db; error = %s", oe.Error())
			}
		}

		return
	})

	return user, err
}
