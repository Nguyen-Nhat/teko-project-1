package setting

type Config struct {
	Server   ServerSetting   `mapstructure:"server"`
	Logger   LoggingSetting  `mapstructure:"logger"`
	Cors     CORSSetting     `mapstructure:"cors"`
	Database DatabaseSetting `mapstructure:"database"`
}
type ServerSetting struct {
	Port           int             `mapstructure:"port"`
	Mode           string          `mapstructure:"mode"`
	MaxRequestBody int64           `mapstructure:"max_request_body"`
	Pageable       PageableSetting `mapstructure:"pageable"`
}
type LoggingSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}
type CORSSetting struct {
	Mode      string          `mapstructure:"mode"`
	WhiteList []CORSWhitelist `mapstructure:"white_list"`
}
type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"allow_origin"`
	AllowHeaders     string `mapstructure:"allow_headers"`
	AllowMethods     string `mapstructure:"allow_methods"`
	ExposeHeaders    string `mapstructure:"expose_headers"`
	AllowCredentials bool   `mapstructure:"allow_credentials"`
}
type DatabaseSetting struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	DbName      string `mapstructure:"dbname"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
}
type PageableSetting struct {
	DefaultPage int32 `mapstructure:"default_page"`
	DefaultSize int32 `mapstructure:"default_size"`
}
