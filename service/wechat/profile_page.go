package wechat

import (
	"project/global"
	"project/model/wechat"
	"project/model/wechat/response"
)

type ProfileService struct{}

// SelectProfileByID 根据用户id查询用户信息
func (p *ProfileService) SelectProfileByID(userId uint) (data response.UserInfo, err error) {
	err = global.TPA_DB.Model(wechat.WxUserInfo{}).Where("id = ?", userId).Find(&data).Error
	return data, err
}
