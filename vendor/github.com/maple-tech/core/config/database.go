package config

// OptionsSQL holds the connection info for the PostgreSQL connection
type OptionsSQL struct {
	//Host for the connection
	Host string `json:"host"`

	//Port for the connections
	Port int `json:"port"`

	//Database name
	Database string `json:"database"`

	//User for login
	User string `json:"user"`

	//Password for login
	Password string `json:"password"`

	//SSL Mode, recommended true, but must be false for HTTP
	SSL bool `json:"ssl"`

	//Create states that the database schema should be whiped and built on startup
	Create bool `json:"cleanOnStart"`

	//Migrate states that the existing schema should be updated
	Migrate bool `json:"migrateOnStart"`
}

// OptionsRedis holds the connection info for the Redis connection
type OptionsRedis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
	PoolSize int    `json:"poolSize"`
	TLS      bool   `json:"tls"`
}

type OptionsStorage struct {
	Azure *AzureStorage `json:"azure"`
}

type AzureStorage struct {
	AccountName      string `json:"account"`
	ContainerName    string `json:"containerName"`
	StorageAccessKey string `json:"storageAccessKey"`
}
