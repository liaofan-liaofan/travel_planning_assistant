package system

import (
	"errors"
	"project/global"
	"project/model/system"
	"project/utils"
)

type UserService struct{}

// Login @author: [chenguanglan](https://github.com/sFFbLL)
// @function: Login
// @description: 用户登录
// @param: u *model.SysUser
// @return: err error, userInter *model.SysUser
func (userService *UserService) Login(u *system.SysUser) (err error, userInter *system.SysUser) {
	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.TPA_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Dept").Preload("Authority").Preload("Authorities.Depts").First(&user).Error
	return err, &user
}

// FindUserByUuid @function: FindUserByUuid
// @description: 通过uuid获取用户信息
// @param: uuid string
// @return: err error, user *model.SysUser
func (userService *UserService) FindUserByUuid(uuid string) (err error, user *system.SysUser) {
	var u system.SysUser
	if err = global.TPA_DB.Preload("Authorities").Preload("Dept").Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
