package response

type UserInfo struct {
	ID         uint   `json:"id"` // 主键ID
	AvatarUrl  string `json:"avatar_url" `
	NickName   string `json:"nick_name" `
	TourismNum uint64 `json:"tourism_num"` // 用户旅游次数
	MapCoupons uint64 `json:"map_coupons"` // 旅游券
}
