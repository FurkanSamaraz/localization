package structures

type ArchiveApp struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	App_id     int    `json:"app_id"`
	Version_id int    `json:"version_id"`
}
type ArchiveMod struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Module_id  int    `json:"module_id"`
	Version_id int    `json:"version_id"`
}
type ArchiveLng struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Language_id int    `json:"language_id"`
	Version_id  int    `json:"version_id"`
}

func (p *ArchiveApp) TableName() string {
	return "localization2.archive_app"
}
func (p *ArchiveMod) TableName() string {
	return "localization2.archive_module"
}
func (p *ArchiveLng) TableName() string {
	return "localization2.archive_lang"
}
