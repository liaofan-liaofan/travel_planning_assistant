package response

import (
	"project/model/wechat"
)

type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type LoginResponse struct {
	User      wechat.WxUserInfo `json:"user"`
	Token     string            `json:"token"`
	ExpiresAt int64             `json:"expiresAt"`
}
