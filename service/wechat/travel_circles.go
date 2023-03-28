package wechat

import (
	"project/global"
	"project/model/wechat"
	"project/model/wechat/request"
	"project/model/wechat/response"
	"strconv"
)

var (
	wxUserTravelCircle wechat.WxUserTravelCircle
	wxTravelCircle     wechat.WxTravelCircle
)

type TravelCirclesService struct{}

func (f *TravelCirclesService) SearchTravelPost(data request.SearchType) (result []response.TravelPost, err error) {
	if data.Page == 0 {
		data.Page = 1
	}
	switch {
	case data.PageSize > 100:
		data.PageSize = 100
	case data.PageSize <= 0:
		data.PageSize = 10
	}
	offset := (data.Page - 1) * data.PageSize

	if data.Type == 1 {
		err = global.TPA_DB.Model(wechat.WxTravelCircle{}).
			Select("wx_user_info.avatar_url,wx_user_info.nick_name,wx_travel_circle.id,wx_travel_circle.user_id,wx_travel_circle.like_num,wx_travel_circle.travel_pictures,wx_travel_circle.article_content,wx_travel_circle.created_at").
			Joins("left join wx_user_info ON wx_travel_circle.user_id = wx_user_info.id").
			Order("like_num desc").Limit(data.PageSize).Offset(offset).Find(&result).Error
	} else if data.Type == 2 {
		err = global.TPA_DB.Model(wechat.WxTravelCircle{}).
			Select("wx_user_info.avatar_url,wx_user_info.nick_name,wx_travel_circle.id,wx_travel_circle.user_id,wx_travel_circle.like_num,wx_travel_circle.travel_pictures,wx_travel_circle.article_content,wx_travel_circle.created_at").
			Joins("left join wx_user_info ON wx_travel_circle.user_id = wx_user_info.id").
			Order("created_at desc").Limit(data.PageSize).Offset(offset).Find(&result).Error
	}
	return result, err
}

// SearchTravelPostAndLikeStatus 查询推荐玩家秀及点赞状态
func (f *TravelCirclesService) SearchTravelPostAndLikeStatus(userId uint, data request.SearchType) (result []response.TravelPostAndLikeStatus, err error) {
	travelPost, err := f.SearchTravelPost(data)
	if err != nil {
		return result, err
	}

	id := make([]uint, len(travelPost))
	var likeStatus []wechat.WxUserTravelCircle
	for i := 0; i < len(travelPost); i++ {
		id[i] = travelPost[i].ID
	}

	if err = global.TPA_DB.Model(wechat.WxUserTravelCircle{}).
		Where("user_id = ? and travel_circle_id in ?", userId, id).
		Find(&likeStatus).Error; err != nil {
		return result, err
	}

	count := make(map[string]int, len(likeStatus))
	for _, v := range likeStatus {
		count[strconv.Itoa(int(v.TravelCircleID))] = v.LikeStatus
	}
	for i := 0; i < len(travelPost); i++ {
		if _, ok := count[strconv.Itoa(int(travelPost[i].ID))]; ok {
			result = append(result, response.TravelPostAndLikeStatus{
				TravelPost: travelPost[i],
				LikeStatus: count[strconv.Itoa(int(travelPost[i].ID))],
			})
		} else {
			result = append(result, response.TravelPostAndLikeStatus{
				TravelPost: travelPost[i],
				LikeStatus: 0,
			})
		}
	}
	return result, err
}

// ConfirmNewPlayerShow 确认发布玩家秀
func (f *TravelCirclesService) ConfirmNewPlayerShow(userId uint, data request.NewPlayShow) (err error) {
	return global.TPA_DB.Create(&wechat.WxTravelCircle{
		UserID:         userId,
		TravelPictures: data.PlayerShowPictures,
		ArticleContent: data.Context,
		TPA_MODEL:      global.TPA_MODEL{},
		//ReleaseTime:        time.Now(),
	}).Error
}
