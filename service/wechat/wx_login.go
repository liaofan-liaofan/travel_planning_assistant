package wechat

import (
	"project/global"
	"project/model/wechat"
	"project/model/wechat/request"
	"project/model/wechat/response"
)

type UserLogin struct{}

var (
	wcUserInfo wechat.WcUserInfo
)

// WXLogin 微信登陆
func (w *UserLogin) WXLogin(wxUser *request.WxCode, wxLoginInfo *response.WXLoginResp) (user *wechat.UserInfo, err error) {
	// 判断用户的OpenID是否存在
	id := wechat.ID{}
	if err = global.TPA_DB.Model(wcUserInfo).Where("open_id = ?", wxLoginInfo.OpenId).Find(&id).Error; err != nil {
		return user, err
	}
	// 不存在则新增
	if id.ID == 0 {
		wxUserInfo := wechat.WcUserInfo{
			Avatar:     wxUser.AvatarUrl,
			UserName:   wxUser.NickName,
			OpenID:     wxLoginInfo.OpenId,
			TourismNum: 0,
			MapCoupons: 0,
		}
		err = global.TPA_DB.Create(&wxUserInfo).Error
	} else { // 存在进行更新用户头像和昵称
		wxUserInfo := wechat.WcUserInfo{
			Avatar:   wxUser.AvatarUrl,
			UserName: wxUser.NickName,
		}
		err = global.TPA_DB.Where("id = ?", id.ID).Updates(wxUserInfo).Error
	}
	return &wechat.UserInfo{ID: id.ID, UserName: wxUser.NickName}, err
}
