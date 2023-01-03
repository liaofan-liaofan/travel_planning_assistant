package config

type Server struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`

	// auto
	AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// redis
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}
