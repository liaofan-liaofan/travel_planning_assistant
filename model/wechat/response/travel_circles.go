package response

import "time"

type User struct {
	AvatarUrl string `json:"avatar_url" `
	NickName  string `json:"nick_name" `
}

// TravelPost 旅游帖
type TravelPost struct {
	User
	ID             uint      `json:"id"`
	UserID         uint      `json:"user_id" gorm:"not null;comment:用户ID"`
	LikeNum        uint64    `json:"like_num" gorm:"not null;comment:文章点赞量"`
	TravelPictures string    `json:"travel_pictures" gorm:"not null;comment:旅游图片"`
	ArticleContent string    `json:"article_content" gorm:"not null;comment:文章内容"`
	CreatedAt      time.Time `json:"created_at"`
}

type TravelPostAndLikeStatus struct {
	TravelPost
	LikeStatus int `json:"like_status"` // 点赞状态(0：未点赞,1：已点赞)
}
