package wechat

import "project/global"

// 旅游圈，晒晒自己旅游的图片和地点

type WcTravelCircles struct {
	global.TPA_MODEL
	ArticleName    string `json:"article_name" gorm:"not null;comment:文章名称"`
	VisitNum       uint64 `json:"visit_num" gorm:"not null;comment:文章访问量"`
	LikeNum        uint64 `json:"like_num" gorm:"not null;comment:文章点赞量"`
	CommentsNum    uint64 `json:"comments_num" gorm:"not null;comment:文章评论量"`
	TravelPictures string `json:"travel_pictures" gorm:"not null;comment:旅游图片"`
	ArticleContent string `json:"article_content" gorm:"not null;comment:文章内容"`
}
