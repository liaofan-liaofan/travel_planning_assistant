package wechat

import "project/global"

// 旅游记录

type WcTravelRecords struct {
	global.TPA_MODEL
	PointSets     string  `json:"point_sets" gorm:"not null;comment:地点集合ID;例如：12,13,15"`
	TotalDistance float64 `json:"total_distance" gorm:"not null;comment:总距离"`
	TotalTourFee  float32 `json:"total_tour_fee" gorm:"not null;comment:总旅游费用"`
}
