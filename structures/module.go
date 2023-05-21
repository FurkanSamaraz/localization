package structures

import "github.com/maple-tech/core/types"

// swagger:parameters Localization_Module
type Module struct {
	// in:body
	// The INT of a thing
	// example: 1
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description types.JSON `json:"description"`
}

func (p *Module) TableName() string {
	return "localization2.module"
}
