package config

// Options holds all the child configuration values needed by core modules
type OptionsQueue struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}
type Options struct {
	Logging OptionsLogging  `json:"logging"`
	Session OptionsSession  `json:"session"`
	Web     OptionsWeb      `json:"web"`
	SQL     OptionsSQL      `json:"database"`
	Redis   OptionsRedis    `json:"redis"`
	Storage *OptionsStorage `json:"storage"`
	Queue   *OptionsQueue   `json:"queue"`
}
