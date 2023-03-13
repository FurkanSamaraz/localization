package types

import "time"

type AppStruct struct {
	Name    string         `json:"name"`
	Alias   string         `json:"alias"`
	Path    string         `json:"path"`
	Modules []ModuleStruct `json:"modules"`
	Version string         `json:"version"`
}

type ModuleStruct struct {
	Name      string           `json:"name"`
	Alias     string           `json:"alias"`
	Path      string           `json:"path"`
	Languages []LanguageStruct `json:"languages"`
}

type LanguageStruct struct {
	Name       string     `json:"name"`
	Alias      string     `json:"alias"`
	Path       string     `json:"path"`
	Version    string     `json:"version"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	BackUpDate *time.Time `json:"backUpDate"`
	DeletedAt  *time.Time `json:"deletedAt"`
	IsBackup   *bool      `json:"isBackup"`
	IsDeleted  *bool      `json:"isDeleted"`
}

/*
	- Retail (App)
		- Checklist
			- English
				- v1.0.0.json

*/
