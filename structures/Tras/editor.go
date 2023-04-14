package structures

import (
	"github.com/maple-tech/core/types"
)

type Editor struct {
	//Application_id Application `json:"application_id"`
	User_id types.User `json:"user_id"`
}

func (p *Editor) TableName() string {
	return "localization.editor"
}
