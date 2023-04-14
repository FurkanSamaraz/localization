package structures

import "github.com/maple-tech/core/types"

type Translation_comment struct {
	Translation_id Translation `json:"translation_id"`
	//Version_id     Version     `json:"version_id"`
	Comment    string     `json:"comment"`
	Created_at string     `json:"created_at"`
	Created_by types.User `json:"created_by"`
}

func (p *Translation_comment) TableName() string {
	return "localization.translation_comment"
}
