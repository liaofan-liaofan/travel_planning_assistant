package wechat

import "project/global"

// 旅游地点集合

type WcLocations struct {
	global.TPA_MODEL
	Pid          uint64 `json:"pid" gorm:"not null;comment:记录父级ID"`
	Name         string `json:"name" gorm:"not null;comment:名称"`
	SubCount     uint   `json:"sub_count"  gorm:"not null;comment:子集数量"`
	DeptSort     uint   `json:"dept_sort" gorm:"not null;comment:排序规则"`
	NameSpelling string `json:"name_spelling" gorm:"not null;comment:名称拼音用于排序"`
	HeatValue    uint   `json:"heat_value" gorm:"not null;comment:热度值"`
}
