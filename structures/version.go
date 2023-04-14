package structures

type Version struct {
	Id       int  `json:"id"`
	Version  int  `json:"version"`
	Is_draft bool `json:"is_draft"`
	Active   bool `json:"active"`
}

func (p *Version) TableName() string {
	return "localization2.version"
}
