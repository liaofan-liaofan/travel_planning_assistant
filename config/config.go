package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`

	// auto
	AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// redis
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	Qiniu Qiniu `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`

	// 前台部分
	WxLogin Wxlogin `mapstructure:"wxlogin" json:"wxlogin" yaml:"wxlogin"`
}
