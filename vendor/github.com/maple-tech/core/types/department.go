package types

import "time"

// Department describes a corporate department
type Department struct {
	ID          ID        `json:"id" db:"d_id"`
	Name        string    `json:"name" db:"d_name"`
	UniqueName  string    `json:"-" db:"d_unique"`
	TimeCreated time.Time `json:"timeCreated" db:"d_created"`
}

// DepartmentMember links a user from the perspective of a department (so user focused, department is implied)
type DepartmentMember struct {
	User

	Supervisor bool      `json:"supervisor" db:"ud_supr"`
	TimeJoined time.Time `json:"timeJoined" db:"ud_created"`
}

// DepartmentComplete describes a corporate department with complete information.
// At this time, it just adds the members of that department
type DepartmentComplete struct {
	Department

	Members []DepartmentMember `json:"members"`
}
