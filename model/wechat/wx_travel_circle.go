package wechat

import "project/global"

// 旅游圈，晒晒自己旅游图片和心得

type WxTravelCircle struct {
	global.TPA_MODEL
	UserID uint `json:"user_id" gorm:"not null;comment:用户ID"`
	//ArticleName    string `json:"article_name" gorm:"not null;comment:文章名称"`
	LikeNum        uint64 `json:"like_num" gorm:"not null;comment:文章点赞量"`
	TravelPictures string `json:"travel_pictures" gorm:"not null;comment:旅游图片"`
	ArticleContent string `json:"article_content" gorm:"not null;comment:文章内容"`
}

func (w WxTravelCircle) TableName() string {
	return "wx_travel_circle"
}
