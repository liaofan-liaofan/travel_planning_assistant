package response

import "project/model/system"

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type AuthorityUser struct {
	Count int64            `json:"count"`
	Users []system.SysUser `json:"users"`
}

type DeptUser struct {
	Id        uint   `json:"user_id"`
	HeaderImg string `json:"url"`
	NickName  string `json:"nick_name"`
}
