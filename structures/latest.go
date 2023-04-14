package structures

type Latest_App struct {
	Id         int    ` json:"id"`
	Name       string `json:"name"`
	App_id     int    `json:"app_id"`
	Version_id int    `json:"version_id"`
}
type Latest_Module struct {
	Id         int    ` json:"id"`
	Name       string `json:"name"`
	Module_id  int    `json:"module_id"`
	Version_id int    `json:"version_id"`
}
type Latest_Lang struct {
	Id         int    ` json:"id"`
	Name       string `json:"name"`
	Lang_id    int    `json:"lang_id"`
	Version_id int    `json:"version_id"`
}

func (p *Latest_App) TableName() string {
	return "localization2.latest_app"
}

func (p *Latest_Module) TableName() string {
	return "localization2.latest_module"
}
func (p *Latest_Lang) TableName() string {
	return "localization2.latest_lang"
}
