package settings

type Config struct {
	Server    ServerSetting   `mapstructure:"server"`
	Databases DatabaseSetting `mapstructure:"databases"`
	Redis     RedisSetting    `mapstructure:"redis"`
	Grafana   GrafanaSetting  `mapstructure:"grafana"`
	Logger    LogSetting      `mapstructure:"logger"`
}
type ServerSetting struct {
	Port int `mapstructure:"port"`
}
type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}
type DatabaseSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	DBName          string `mapstructure:"dbname"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	MaxIdleConns    int    `mapstructure:"maxIdleconns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}
type GrafanaSetting struct {
	URL      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
}
type LogSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_size      int    `mapstructure:"max_size"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Compress      bool   `mapstructure:"compress"`
}
