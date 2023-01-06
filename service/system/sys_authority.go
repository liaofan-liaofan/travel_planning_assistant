package system

import (
	"errors"
	"gorm.io/gorm"
	"project/global"
	"project/model/common/request"
	"project/model/system"
	"strconv"
)

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

// CreateAuthority @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: CreateAuthority
// @description: 创建一个角色
// @param: auth model.SysAuthority
// @return: err error, authority model.SysAuthority
func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (err error, authority system.SysAuthority) {
	err = global.TPA_DB.Create(&auth).Error
	return err, auth
}

// DeleteAuthority @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: DeleteAuthority
// @description: 删除角色
// @param: auth *model.SysAuthority
// @return: err error
func (authorityService *AuthorityService) DeleteAuthority(auth *system.SysAuthority) (err error) {
	if !errors.Is(global.TPA_DB.Where("sys_authority_authority_id = ?", auth.AuthorityId).First(&system.SysUseAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	return global.TPA_DB.Transaction(func(tx *gorm.DB) error {
		db := global.TPA_DB.Preload("SysBaseMenus").Preload("Depts").Where("authority_id = ?", auth.AuthorityId).First(auth)
		if err := db.Select("SysBaseMenus").Select("Depts").Delete(auth).Error; err != nil {
			return err
		}
		if success := CasbinServiceApp.ClearCasbin(0, strconv.Itoa(int(auth.AuthorityId))); !success {
			return nil
		}
		return nil
	})
}

// GetAuthorityInfoList @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: GetAuthorityInfoList
// @description: 分页获取数据
// @param: info request.PageInfo
// @return: err error, list interface{}, total int64
func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var authority []system.SysAuthority
	err = global.TPA_DB.Limit(limit).Offset(offset).Preload("Depts").Find(&authority).Error
	return err, authority, total
}

// UpdateAuthority @author: [chenguanglan](https://github.com/sFFbLL)
// @function: UpdateAuthority
// @description: 更改一个角色
// @param: auth model.SysAuthority
// @return: err error, authority model.SysAuthority
func (authorityService *AuthorityService) UpdateAuthority(auth system.SysAuthority, deptId []uint) (err error, authority system.SysAuthority) {
	depts := make([]system.SysDept, 0)
	if auth.DataScope == "自定义" {
		for _, d := range deptId {
			depts = append(depts, system.SysDept{TPA_MODEL: global.TPA_MODEL{ID: d}})
		}
	}
	err = global.TPA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&auth).Association("Depts").Replace(depts)
		if err != nil {
			return err
		}
		return tx.Where("authority_id = ?", auth.AuthorityId).Preload("Depts").First(&system.SysAuthority{}).Updates(&auth).Error
	})
	return err, auth
}

// GetAuthorityInfo @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: GetAuthorityInfoByIDs
// @description: 根据角色id切片获取所有角色信息
// @param: authIds []uint
// @return: err error, sa []model.SysAuthority

func (authorityService *AuthorityService) GetAuthorityInfoByIDs(authIds []uint) (err error, sa []system.SysAuthority) {
	err = global.TPA_DB.Preload("SysBaseMenus").Preload("Depts").Where("authority_id in (?)", authIds).Find(&sa).Error
	return err, sa
}

// GetAuthorityBasicInfo @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: GetAuthorityBasicInfo
// @description: 获取基本角色信息
// @param: auth model.SysAuthority
// @return: err error, sa model.SysAuthority

func (authorityService *AuthorityService) GetAuthorityBasicInfo(auth system.SysAuthority) (err error, sa system.SysAuthority) {
	err = global.TPA_DB.Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

// SetMenuAuthority @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: SetMenuAuthority
// @description: 菜单与角色绑定
// @param: auth *model.SysAuthority
// @return: error

func (authorityService *AuthorityService) SetMenuAuthority(auth *system.SysAuthority) error {
	var s system.SysAuthority
	global.TPA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.TPA_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}
