package structures

import "github.com/maple-tech/core/types"

type Module struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description types.JSON `json:"description"`
}

func (p *Module) TableName() string {
	return "localization2.module"
}
