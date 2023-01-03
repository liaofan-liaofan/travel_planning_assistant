package config

type Server struct {
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`

	// auto
	AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
