package structures

type Translation struct {
	Id int `json:"id"`
	//	Version_id Version `json:"version_id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Create_at string `json:"create_at"`
}

func (p *Translation) TableName() string {
	return "localization.translation"
}
