package wechat

import (
	"project/global"
)

// 收藏夹

type WxFavorite struct {
	global.TPA_MODEL
	UserID    uint   `json:"user_id" gorm:"not null;comment:用户ID"`
	TourName  string `json:"tour_name" gorm:"not null;comment:旅游规划的名称;默认为新启航"`
	PointSets string `json:"point_sets" gorm:"not null;comment:地点集合"`
}

func (w WxFavorite) TableName() string {
	return "wx_favorite"
}
