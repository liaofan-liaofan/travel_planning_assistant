package wechat

import "project/global"

type WxUserTravelCircle struct {
	global.TPA_MODEL
	UserID         uint `json:"user_id" gorm:"not null;comment:用户ID"`
	TravelCircleID uint `json:"travel_circle_id" gorm:"not null;comment:旅游帖ID"`
	LikeStatus     int  `json:"like_status" gorm:"not null;comment:点赞状态(0：未点赞,1：已点赞)"`
}

func (w WxUserTravelCircle) TableName() string {
	return "wx_user_travel_circle"
}
