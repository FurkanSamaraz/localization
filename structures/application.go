package structures

import "github.com/maple-tech/core/types"

// swagger:parameters Localization_App
// swagger:parameters Localization_Apps

type Application struct {
	// in:body
	// The INT of a thing
	// example: 1
	Id int ` json:"id"`
	// The Name of a thing
	// example: Some name
	Name string `json:"name"`
	// The Description of a thing
	// example: Some description
	Description types.JSON `json:"description"`
} // @name ThingsResponse

func (p *Application) TableName() string {
	return "localization2.app"
}
