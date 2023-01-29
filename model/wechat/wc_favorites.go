package wechat

import (
	"project/global"
	"time"
)

// 收藏夹

type WcFavorites struct {
	global.TPA_MODEL
	TourName         string    `json:"tour_name" gorm:"not null;comment:旅游规划的名称;默认为新启航"`
	StartTime        time.Time `json:"start_time" gorm:"not null;comment:旅游开始日期时间"`
	PointSets        string    `json:"point_sets" gorm:"not null;comment:地点集合ID;例如：12,13,15"`
	TravelTime       float32   `json:"travel_time" gorm:"not null;comment:路程时间;例如：20.5小时"`
	DistanceTraveled float64   `json:"distance_traveled" gorm:"not null;comment:路程距离"`
	DistanceCosts    float32   `json:"distance_costs" gorm:"not null;comment:路程费用估计"`
}
