package setting

type Config struct {
	Server ServerSetting  `mapstructure:"server"`
	Logger LoggingSetting `mapstructure:"logger"`
}
type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}
type LoggingSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}
