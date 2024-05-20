package configs

type ServerConfig struct {
	System    System    `mapstructure:"system" yaml:"system"`
	LogConfig LogConfig `mapstructure:"log" yaml:"log"`
}
