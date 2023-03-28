package request

import "project/model/common/request"

type SearchType struct {
	Type int `json:"type" form:"type" binding:"required"` // 查询类型1：推荐; 2:最新
	request.PageInfo
}

type NewPlayShow struct {
	PlayerShowPictures string `json:"player_show_pictures" form:"player_show_pictures" binding:"required"` // 玩家秀图片地址ID
	Context            string `json:"context" form:"player_show_pictures"`
}
