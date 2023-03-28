package wechat

import (
	uuid "github.com/satori/go.uuid"
	"project/global"
)

// 用户信息表

type WxUserInfo struct {
	global.TPA_MODEL
	OpenID string    `json:"open_id" gorm:"not null;index:idx_wx_user_info_id;index:idx_wx_user_info_open_id;comment:小程序唯一标识"`
	UUID   uuid.UUID `json:"uuid" gorm:"not null;index:idx_wx_user_info_id;index:idx_wx_user_info_uu_id;comment:用户UUID"`
	//UserID     uint64 `json:"user_id" gorm:"not null;comment:用户ID"`
	AvatarUrl  string `json:"avatar_url" gorm:"not null;comment:用户头像"`
	NickName   string `json:"nick_name"  gorm:"not null;comment:用户昵称"`
	TourismNum uint64 `json:"tourism_num" gorm:"not null;comment:用户旅游次数"`
	MapCoupons uint64 `json:"map_coupons" gorm:"not null;comment:旅游券"`
}

type ID struct {
	ID   uint      // 主键ID
	UUID uuid.UUID `json:"uuid" gorm:"not null;comment:用户UUID"`
}

func (w WxUserInfo) TableName() string {
	return "wx_user_info"
}
