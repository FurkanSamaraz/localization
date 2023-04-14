package types

import "time"

//Position describes a corporate position for an employee
type Position struct {
	ID          ID        `json:"id" db:"p_id"`
	Name        string    `json:"name" db:"p_name"`
	UniqueName  string    `json:"-" db:"p_unique"`
	TimeCreated time.Time `json:"timeCreated" db:"p_created"`
}

//PositionMember links a user from the perspective of a position
type PositionMember struct {
	User

	TimeJoined time.Time `json:"timeJoined" db:"up_created"`
}

//PositionComplete wraps the Postion object while including the members of it
type PositionComplete struct {
	Position

	Members []PositionMember `json:"members"`
}
