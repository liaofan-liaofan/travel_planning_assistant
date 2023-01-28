package config

// Wxlogin 微信登录
type Wxlogin struct {
	GrantType string `mapstructure:"grantype" json:"grantype" yaml:"grantype"` // 授权类型
	AppId     string `mapstructure:"appid" json:"appid" yaml:"appid"`          // 小程序ID
	Secret    string `mapstructure:"secret" json:"secret" yaml:"secret"`       // 小程序密匙
}
