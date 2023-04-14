package structures

import (
	"github.com/maple-tech/core/types"
)

type Translation_history struct {
	//	Version_id Version    `json:"version_id"`
	Key        string     `json:"key"`
	Value      string     `json:"value"`
	Create_at  string     `json:"create_at"`
	Created_by types.User `json:"created_by"`
}

func (p *Translation_history) TableName() string {
	return "localization.translation_history"
}
