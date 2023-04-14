package structures

type TrashApp struct {
	Id         int    ` json:"id"`
	Name       string `json:"name"`
	App_id     int    `json:"app_id"`
	Version_id int    `json:"version_id"`
}

type TrashMod struct {
	Id         int    ` json:"id"`
	Name       string `json:"name"`
	Module_id  int    `json:"module_id"`
	Version_id int    `json:"version_id"`
}
type TrashLang struct {
	Id          int    ` json:"id"`
	Name        string `json:"name"`
	Language_id int    `json:"language_id"`
	Version_id  int    `json:"version_id"`
}

func (p *TrashApp) TableName() string {
	return "localization2.trash_app"
}

func (p *TrashMod) TableName() string {
	return "localization2.trash_module"
}

func (p *TrashLang) TableName() string {
	return "localization2.trash_lang"
}
