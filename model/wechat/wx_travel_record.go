package wechat

import "project/global"

// 旅游记录

type WxTravelRecord struct {
	global.TPA_MODEL
	UserID        uint    `json:"user_id" gorm:"not null;comment:用户ID"`
	PointSets     string  `json:"point_sets" gorm:"not null;comment:地点集合ID;例如：12,13,15"`
	TotalDistance float64 `json:"total_distance" gorm:"not null;comment:总距离"`
	TotalTourFee  float32 `json:"total_tour_fee" gorm:"not null;comment:总旅游费用"`
}

func (w WxTravelRecord) TableName() string {
	return "wx_travel_record"
}
