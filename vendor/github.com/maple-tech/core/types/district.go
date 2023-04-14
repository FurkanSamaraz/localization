package types

import "time"

//District holds the information labelling a district within a company.
//These are tied to a country, and a region as well.
type District struct {
	ID        ID          `json:"id" db:"dis_id"`
	CountryID CountryCode `json:"country" db:"dis_cnt"`
	RegionID  RegionCode  `json:"region" db:"dis_rgn"`

	Name        string    `json:"name" db:"dis_name"`
	TimeCreated time.Time `json:"timeCreated" db:"dis_created"`
	TimeUpdated NullTime  `json:"timeUpdated" db:"dis_updated"`
}
