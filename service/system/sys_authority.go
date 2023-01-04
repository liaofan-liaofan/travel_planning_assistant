package system

import (
	"project/global"
	"project/model/system"
)

type AuthorityService struct{}

// GetAuthorityInfo @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: GetAuthorityInfoByIDs
// @description: 根据角色id切片获取所有角色信息
// @param: authIds []uint
// @return: err error, sa []model.SysAuthority

func (authorityService *AuthorityService) GetAuthorityInfoByIDs(authIds []uint) (err error, sa []system.SysAuthority) {
	err = global.TPA_DB.Preload("SysBaseMenus").Preload("Depts").Where("authority_id in (?)", authIds).Find(&sa).Error
	return err, sa
}
