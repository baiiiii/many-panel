package configs

type LogConfig struct {
	Level     string `mapstructure:"level" yaml:"level"`
	TimeZone  string `mapstructure:"timeZone" yaml:"timeZone"`
	LogName   string `mapstructure:"log_name" yaml:"log_name"`
	LogSuffix string `mapstructure:"log_suffix" yaml:"log_suffix"`
	MaxBackup int    `mapstructure:"max_backup" yaml:"max_backup"`
}
