package types

import "time"

//Location is an arbitrary location within a company,
//these can be anything from meeting places, carpool lots,
//route stops, etc.
type Location struct {
	ID         ID          `json:"id" db:"loc_id"`
	CountryID  CountryCode `json:"country" db:"loc_cnt"`
	RegionID   RegionCode  `json:"region" db:"loc_rgn"`
	DistrictID ID          `json:"district" db:"loc_dist"`

	Name string `json:"name" db:"loc_name"`

	//TODO: GPS points and polygons

	TimeCreated time.Time `json:"timeCreated" db:"loc_created"`
	TimeUpdated NullTime  `json:"timeUpdated" db:"loc_updated"`
}
