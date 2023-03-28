package wechat

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"project/global"
	"project/model/wechat"
	"project/model/wechat/request"
	"project/model/wechat/response"
)

type UserLoginService struct{}

var (
	wxUserInfo wechat.WxUserInfo
)

// WXLogin 微信登陆
func (w *UserLoginService) WXLogin(userinfo request.WxCode, wxLoginInfo *response.WXLoginResp) (user wechat.WxUserInfo, err error) {
	// 判断用户的OpenID是否存在
	if err = global.TPA_DB.Model(wxUserInfo).Where("open_id = ?", wxLoginInfo.OpenId).First(&user).Error; err != nil {
		user = wechat.WxUserInfo{
			UUID:       uuid.NewV4(),
			OpenID:     wxLoginInfo.OpenId,
			NickName:   userinfo.NickName,
			AvatarUrl:  userinfo.AvatarUrl,
			TourismNum: 0,
			MapCoupons: 0,
		}
		err = global.TPA_DB.Create(&user).Error
	} else {
		if user.AvatarUrl != userinfo.AvatarUrl && user.NickName != userinfo.NickName {
			err = global.TPA_DB.Where("id = ?", user.ID).
				Update("avatar", userinfo.AvatarUrl).Update("user_name", userinfo.NickName).Error
			user.AvatarUrl, user.NickName = userinfo.AvatarUrl, userinfo.NickName
		} else if user.AvatarUrl != userinfo.AvatarUrl {
			err = global.TPA_DB.Where("id = ?", user.ID).
				Update("avatar", userinfo.AvatarUrl).Error
			user.AvatarUrl = userinfo.AvatarUrl
		} else if user.NickName != userinfo.NickName {
			err = global.TPA_DB.Where("id = ?", user.ID).
				Update("user_name", userinfo.NickName).Error
			user.NickName = userinfo.NickName
		}
	}
	return user, err
}

// FindUserByUuid @function: FindUserByUuid
// @description: 通过uuid获取用户信息
// @param: uuid string
// @return: err error, user *model.WcUserInfo
func (w *UserLoginService) FindUserByUuid(uuid string) (err error, user *wechat.WxUserInfo) {
	var u wechat.WxUserInfo
	if err = global.TPA_DB.Model(u).Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
