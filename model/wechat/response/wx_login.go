package response

import (
	"project/model/wechat"
)

// LoginData 返回用户登录详细列表
type LoginData struct {
	Token string    `json:"token" `
	User  LoginUser `json:"user"`
}

type LoginUser struct {
	Id        int64  `json:"id" binding:"required"`         // 用户id
	NickName  string `json:"nick_name" binding:"required"`  // 用户昵称
	AvatarUrl string `json:"avatar_url" binding:"required"` // 用户头像
}

type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type LoginResponse struct {
	User      wechat.UserInfo `json:"user"`
	Token     string          `json:"token"`
	ExpiresAt int64           `json:"expiresAt"`
}
