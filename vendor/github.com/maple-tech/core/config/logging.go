package config

//OptionsLogging holds the configuration options for the logging portion of the service
type OptionsLogging struct {
	//Level handles the logging level specifics, use the enums {DEBUG|INFO|ERROR}
	Level string `json:"level"`

	//Path fullyqualified path to the output file
	Path string `json:"path"`

	//Format set's the formatter the internal Logrus library will use, accepts {TEXT|JSON}
	Format string `json:"format"`

	//Rotate specifies whether we should bind and use the log rotator provided by logrus
	Rotate bool `json:"rotate"`

	//MaxSize is the maximum file size in megabytes before the log file will rotate
	MaxSize int `json:"max-size"`

	//MaxAge is the maximum date age in days before the log file will rotate
	MaxAge int `json:"max-age"`

	//Backups holds the maximum number of backed up logs we should keep
	Backups int `json:"backups"`

	//Compress flags whether backups should be compressed (gzip)
	Compress bool `json:"compress"`

	//Colors sets whether the output (when using TEXT format) should include color codes
	Colors bool `json:"colored"`
}