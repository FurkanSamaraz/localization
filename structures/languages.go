package structures

import "github.com/maple-tech/core/types"

type Language struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description types.JSON `json:"description"`
}

// type Language struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	UniqeName   string `json:"uniqe_name"`
// 	Create_at   string `json:"create_at"`
// }

func (p *Language) TableName() string {
	return "localization2.language"
}
