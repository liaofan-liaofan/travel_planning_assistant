package system

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"project/global"
	"project/model/common/request"
	"project/model/system"
	"project/utils"
	"strconv"
	"strings"
)

type UserService struct{}

// Login @author: [liaofan](https://github.com/liaofan-liaofan)
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

// Register @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: Register
// @description: 用户注册
// @param: u model.SysUser
// @return: err error, userInter model.SysUser

func (userService *UserService) Register(u system.SysUser) (err error, userInter system.SysUser) {
	var user system.SysUser
	if !errors.Is(global.TPA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.TPA_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Create(&u).Error
		if TxErr != nil {
			return TxErr
		}
		return nil
	})
	return err, u
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info request.PageInfo, deptId []uint, isAll bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.TPA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	if !isAll {
		db = db.Where("dept_id in (?)", deptId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Dept").Find(&userList).Error
	return err, userList, total
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
// FindUserInfoByAuthority @function: FindUserInfoByAuthority
// @description: 通过角色id获取用户信息
// @param: uuid string
// @return: err error, user []uint

func (userService *UserService) FindUserInfoByAuthority(authorityId uint) (err error, user []system.SysUser, count int64) {
	userId := make([]uint, 0)
	err = global.TPA_DB.Model(system.SysUseAuthority{}).Distinct("sys_user_id").Where("`sys_authority_authority_id` = ?", authorityId).Find(&userId).Error
	err = global.TPA_DB.Where("id in (?)", userId).Find(&user).Count(&count).Error
	return
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: UpdatePassword
//@description: 用户本人修改密码
//@param: user *system.SysUser, newPassword string
//@return: err error, sysUser *system.SysUser

func (userService *UserService) UpdatePassword(u *system.SysUser, newPassword string) (err error, sysUser *system.SysUser) {
	var user *system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.TPA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: ResetPassword
//@description: 重置密码
//@param: user *system.SysUser, newPassword string
//@return: err error, sysUser *system.SysUser

func (userService *UserService) ResetPassword(u *system.SysUser, newPassword string) (err error, sysUser *system.SysUser) {
	var user *system.SysUser
	err = global.TPA_DB.Where("id", u.ID).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

// @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: GetUserInfo
// @description: 获取用户信息
// @param: uuid uuid.UUID
// @return: err error, user system.SysUser

func (userService *UserService) GetUserInfo(uuid string) (err error, user system.SysUser) {
	var reqUser system.SysUser
	err = global.TPA_DB.Preload("Authorities").Preload("Authority").Preload("Dept").First(&reqUser, "uuid = ?", uuid).Error
	return err, reqUser
}

// FindUserById @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: FindUserById
// @description: 通过id获取用户信息
// @param: id int
// @return: err error, user *model.SysUser

func (userService *UserService) FindUserById(id uint) (err error, user *system.SysUser) {
	var u system.SysUser
	err = global.TPA_DB.Preload("Authorities").Preload("Dept").Where("`id` = ?", id).First(&u).Error
	return err, &u
}

// SetUserAuthorities @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: SetUserAuthorities
// @description: 设置一个用户的权限
// @param: id uint, authorityIds []uint
// @return: err error
func (userService *UserService) SetUserAuthorities(updateUser system.SysUser, authorityIds []uint) (err error) {
	return global.TPA_DB.Transaction(func(tx *gorm.DB) error {
		useAuthority := make([]system.SysAuthority, 0)
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.SysAuthority{
				AuthorityId: v,
			})
		}
		TxErr := tx.Model(&updateUser).Association("Authorities").Replace(useAuthority)
		if TxErr != nil {
			return TxErr
		}
		//更新缓存
		authorityIdJson, _ := json.Marshal(authorityIds)
		err = global.TPA_REDIS.HSet(context.Background(), updateUser.UUID.String(), "authorityId", authorityIdJson).Err()
		if err != nil {
			return err
		}
		return nil
	})
}

// @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: SetUserInfo
// @description: 修改用户信息
// @param: reqUser model.SysUser
// @return: err error, user model.SysUser

func (userService *UserService) SetUserInfo(reqUser system.SysUser) (err error, user system.SysUser) {
	tx := global.TPA_DB.Begin()
	err = tx.Updates(&reqUser).Error
	if err != nil {
		tx.Rollback()
		return
	}
	userInfo := reqUser
	//更新deptId
	userInfo.DeptId = reqUser.DeptId
	err = global.TPA_REDIS.HSet(context.Background(), reqUser.UUID.String(), "deptId", reqUser.DeptId).Err()
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return err, reqUser
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: SetUserDept
//@description: 修改用户部门
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetUserDept(reqUser system.SysUser) (err error, user system.SysUser) {
	tx := global.TPA_DB.Begin()
	err = tx.Updates(&reqUser).Error
	if err != nil {
		tx.Rollback()
		return
	}
	userInfo := reqUser
	//更新deptId
	userInfo.DeptId = reqUser.DeptId
	err = global.TPA_REDIS.HSet(context.Background(), reqUser.UUID.String(), "deptId", reqUser.DeptId).Err()
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return err, reqUser
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: SetSelfInfo
//@description: 修改用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetSelfInfo(reqUser system.SysUser) (err error, user system.SysUser) {
	err = global.TPA_DB.Updates(&reqUser).Error
	return err, reqUser
}

// DeleteUser @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: DeleteUser
// @description: 删除用户
// @param: id uint
// @return: err error

func (userService *UserService) DeleteUser(id uint) (err error) {
	return global.TPA_DB.Transaction(func(tx *gorm.DB) error {
		var user system.SysUser
		TxErr := global.TPA_DB.Delete(&user, id).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = global.TPA_DB.Delete(&system.SysUseAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		//TxErr = CasbinServiceApp.DeleteUserAuthority(id)
		//if TxErr != nil {
		//	return tx.Rollback().Error
		//}
		return nil
	})
}

// @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: ParseExcelToDataList
// @description: 加载excel
// @return: user []system.SysUser, err error

func (userService *UserService) ParseExcelToDataList() ([]system.SysUser, error) {
	skipHeader := true
	fixedHeader := []string{"ID", "用户名", "昵称", "手机号", "邮箱", "角色名称", "部门名称"}
	file, err := excelize.OpenFile(global.TPA_CONFIG.Excel.Dir + "ExcelImport.xlsx")
	if err != nil {
		return nil, err
	}
	users := make([]system.SysUser, 0)
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		if skipHeader {
			if utils.CompareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return nil, errors.New("excel格式错误")
			}
		}
		if len(row) != len(fixedHeader) {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		contains := strings.Contains(row[5], "、")
		if contains {
			split := strings.Split(row[5], "、")
			var authorities []system.SysAuthority
			for i := 0; i < len(split); i++ {
				authorities = append(authorities, system.SysAuthority{AuthorityName: split[i]})
			}
			user := system.SysUser{
				TPA_MODEL: global.TPA_MODEL{
					ID: uint(id),
				},
				Username:    row[1],
				NickName:    row[2],
				Phone:       row[3],
				Email:       row[4],
				Authorities: authorities,
				Dept:        system.SysDept{DeptName: row[6]},
			}
			users = append(users, user)
		} else {
			user := system.SysUser{
				TPA_MODEL: global.TPA_MODEL{
					ID: uint(id),
				},
				Username:  row[1],
				NickName:  row[2],
				Phone:     row[3],
				Email:     row[4],
				Authority: system.SysAuthority{AuthorityName: row[5]},
				Dept:      system.SysDept{DeptName: row[6]},
			}
			users = append(users, user)
		}

	}
	return users, nil
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: ParseDataListToExcel
//@description: 导出excel
//@param: info []system.SysUser, path string
//@return: err error

func (userService *UserService) ParseDataListToExcel(info []system.SysUser, path string) error {
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "用户名", "昵称", "手机号", "邮箱", "角色名称", "部门名称"})
	for i, user := range info {
		axis := fmt.Sprintf("A%d", i+2)
		userAuthority := ""
		for index, authority := range user.Authorities {
			if index != len(user.Authorities)-1 {
				userAuthority += authority.AuthorityName + ", "
			}
			userAuthority += authority.AuthorityName
		}
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			user.ID,
			user.Username,
			user.NickName,
			user.Phone,
			user.Email,
			userAuthority,
			user.Dept.DeptName,
		})
	}
	mkdirErr := os.MkdirAll(global.TPA_CONFIG.Excel.Dir, os.ModePerm)
	if mkdirErr != nil {
		global.TPA_LG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	err := excel.SaveAs(path)
	return err
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: Template
//@description: 下载模板
//@param: path string
//@return: err error

func (userService *UserService) Template(path string) error {
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "用户名", "昵称", "手机号", "邮箱", "角色名称", "部门名称"})
	axis := fmt.Sprintf("A%d", 2)
	excel.SetSheetRow("Sheet1", axis, &[]interface{}{
		"1",
		"test",
		"测试用户",
		"156***",
		"xxx@xx.com",
		"超级管理员",
		"顶级部门",
	})
	axisT := fmt.Sprintf("A%d", 3)
	excel.SetSheetRow("Sheet1", axisT, &[]interface{}{
		"1",
		"test",
		"测试用户",
		"156***",
		"xxx@xx.com",
		"超级管理员、测试管理员",
		"顶级部门",
	})
	err := excel.SaveAs(path)
	return err
}
