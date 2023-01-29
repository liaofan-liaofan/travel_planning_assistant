package wechat

import "project/global"

// 用户信息表

type WcUserInfo struct {
	global.TPA_MODEL
	//UserID     uint64 `json:"user_id" gorm:"not null;comment:用户ID"`
	Avatar     string `json:"avatar" gorm:"not null;comment:用户头像"`
	UserName   string `json:"user_name"  gorm:"not null;comment:用户昵称"`
	OpenID     string `json:"open_id" gorm:"not null;comment:小程序唯一标识"`
	TourismNum uint64 `json:"tourism_num" gorm:"not null;comment:用户旅游次数"`
	MapCoupons uint64 `json:"map_coupons" gorm:"not null;comment:旅游计算券"`
}
