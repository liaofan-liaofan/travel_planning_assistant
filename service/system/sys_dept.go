package system

import (
	"errors"
	"gorm.io/gorm"
	"project/global"
	"project/model/common/request"
	"project/model/system"
	"project/model/system/response"
)

type DeptService struct{}

// AddDepartment @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: AddDepartment
// @description: 新增部门
// @param: dept system.SysDept
// @return: error
func (departService *DeptService) AddDepartment(dept system.SysDept) (err error, sysDept system.SysDept) {
	if !errors.Is(global.TPA_DB.Where("dept_name = ?", dept.DeptName).First(&system.SysDept{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name"), dept
	}
	err = global.TPA_DB.Create(&dept).Error
	return err, dept
}

// DeleteDepartment @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: DeleteDepartment
// @description: 删除部门
// @param: id uint
// @return: error
func (departService *DeptService) DeleteDepartment(dept *system.SysDept) (err error) {
	if errors.Is(global.TPA_DB.First(&dept).Error, gorm.ErrRecordNotFound) {
		return errors.New("该部门不存在")
	}
	if !errors.Is(global.TPA_DB.Where("parent_id = ?", dept.ID).First(&system.SysDept{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此部门下有子部门禁止删除")
	}
	if !errors.Is(global.TPA_DB.Where("dept_id = ?", dept.ID).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此部门下有用户禁止删除")
	}
	err = global.TPA_DB.Where("parent_id = ?", dept.ID).First(&system.SysDept{}).Error
	if err != nil {
		err = global.TPA_DB.Where("id = ?", dept.ID).Delete(&dept).Error
		if err != nil {
			return err
		}
	} else {
		return errors.New("存在子部门，不能删除")
	}
	return err
}

// UpdateDepartment @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: UpdateDepartment
// @description: 修改部门
// @param: dept system.SysDept
// @return: error
func (departService *DeptService) UpdateDepartment(dept system.SysDept) (err error, sysDept system.SysDept) {
	err = global.TPA_DB.Where("id = ?", dept.ID).First(&system.SysDept{}).Updates(&dept).Error
	return err, dept
}

// GetDeptList @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: GetDeptList
// @description: 查询部门列表
// @param: info request.PageInfo
// @return: error
func (departService *DeptService) GetDeptList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.TPA_DB.Model(&system.SysDept{})
	var dept []system.SysDept
	err = db.Where("parent_id = ?", 0).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&dept).Order("dept_sort").Error
	if len(dept) > 0 {
		for i := range dept {
			err = departService.findChildrenDepartment(&dept[i])
		}
	}
	return err, dept, total
}

// GetDeptListById @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: GetDeptListById
// @description: 根据pid查询部门列表
// @param: id uint
// @return: error
func (departService *DeptService) GetDeptListById(id uint) (err error, list interface{}, total int64) {
	db := global.TPA_DB.Model(&system.SysDept{})
	var dept []system.SysDept
	err = db.Where("parent_id = ?", id).Find(&dept).Order("dept_sort").Error
	err = db.Where("parent_id = ? ", id).Count(&total).Error
	return err, dept, total
}

// findChildrenDepartment @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: findChildrenDepartment
// @description: 查询子菜单
// @param: dept *system.SysDept
// @return: err error
func (departService *DeptService) findChildrenDepartment(dept *system.SysDept) (err error) {
	err = global.TPA_DB.Where("parent_id = ?", dept.ID).Find(&dept.Children).Error
	if len(dept.Children) > 0 {
		for i := range dept.Children {
			err = departService.findChildrenDepartment(&dept.Children[i])
		}
	}
	return err
}

// GetUserByDeptId @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: GetUserByDeptId
// @description: 获取部门下的用户
// @param: dept system.SysDept
// @return: err error, user []system.SysDept
func (departService *DeptService) GetUserByDeptId(dept system.SysDept) (err error, userRes []interface{}) {
	var userDept response.GetUserByDeptId
	var user []response.DeptUser
	var count int64
	err = global.TPA_DB.Table("sys_depts").Select("id, dept_name").Where("id = ?", dept.ID).Find(&userDept).Error
	if err != nil {
		return err, nil
	}
	err = global.TPA_DB.Table("sys_users").Select("id, nick_name,header_img").Where("dept_id = ?", dept.ID).Find(&user).Count(&count).Error
	userDept.Count = int(count)
	userRes = append(userRes, userDept)
	for i := 0; i < userDept.Count; i++ {
		userRes = append(userRes, user[i])
	}
	return err, userRes
}
