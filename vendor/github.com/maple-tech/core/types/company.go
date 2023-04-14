package types

import (
	"time"

	sqltypes "github.com/jmoiron/sqlx/types"
)

// Company holds the simple information about a company
type Company struct {
	ID         URID    `json:"id" db:"c_id"`
	Name       string  `json:"name" db:"c_name"`
	UniqueName string  `json:"uniqueName" db:"c_unique"`
	Picture    *string `json:"picture,omitempty" db:"c_picture"`
}

// CompanyInfo contains public information (profile) for a company
type CompanyInfo struct {
	InfoID      ID                `json:"infoId" db:"ci_id"`
	Employees   int               `json:"employees" db:"ci_emp"`
	Industries  IndustryArray     `json:"industries" db:"ci_ind"`
	ContactInfo sqltypes.JSONText `json:"contactInfo" db:"ci_contact"`
	TimeUpdated NullTime          `json:"timeUpdated,omitempty" db:"ci_updated"`
	// TimeUpdated NullTime `json:"-,omitempty" db:"ci_updated"`
}

// CompanyModule subscription to a specific module, or "service" of Maple. Includes end-of-period for payment
type CompanyModule struct {
	ID          ID                `json:"-" db:"cm_id"`
	Module      string            `json:"module" db:"cm_mod"`
	Settings    sqltypes.JSONText `json:"settings" db:"cm_settings"`
	Active      bool              `json:"active" db:"cm_active"`
	EndOfPeriod time.Time         `json:"endOfPeriod" db:"cm_eop"`
	TimeCreated time.Time         `json:"timeCreated" db:"cm_created"`
	TimeUpdated NullTime          `json:"timeUpdated,omitempty" db:"cm_updated"`
}

// CompanyRecord binds together a basic "company" and it's information
type CompanyRecord struct {
	Company
	CompanyInfo
}

// CompanyMember describes a user within a company. It is different then UserCompany
// in that this is from the perspective of the company and thus does not duplicate the CompanyRecord
type CompanyMember struct {
	User

	Departments []UserDepartment `json:"departments"`
	Positions   []UserPosition   `json:"positions"`

	MembershipActive      bool      `json:"membershipActive" db:"uc_active"`
	TimeJoined            time.Time `json:"timeJoined" db:"uc_joined"`
	TimeMembershipUpdated NullTime  `json:"timeMembershipUpdated,omitempty" db:"uc_updated"`

	Role   PermissionModal `json:"role"`
	RoleID ID              `json:"-" db:"uc_role"`

	IsExternal             bool     `json:"isExternal" db:"uc_external"`
	ExternalCompany        *Company `json:"externalCompany"`
	ExternalCompanyID      *URID    `json:"-" db:"uce_id"`
	ExternalCompanyName    *string  `json:"-" db:"uce_name"`
	ExternalCompanyPicture *string  `json:"-" db:"uce_picture"`
}

// CompanyComplete holds ALL information (pertenent) for a company
type CompanyComplete struct {
	CompanyRecord

	Modules []CompanyModule `json:"modules"`

	Departments []Department `json:"departments"`
	Positions   []Position   `json:"positions"`

	Roles []PermissionModal `json:"roles"`

	Members []CompanyMember `json:"members"`
}
