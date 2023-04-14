package structures

import "github.com/maple-tech/core/types"

type Application struct {
	Id          int        ` json:"id"`
	Name        string     `json:"name"`
	Description types.JSON `json:"description"`
}

func (p *Application) TableName() string {
	return "localization2.app"
}
