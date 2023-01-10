package system

import (
	"errors"
	"gorm.io/gorm"
	"project/global"
	"project/model/common/request"
	"project/model/system"
)

type ApiService struct{}

// ApiService @author: [liaofan](https://github.com/liaofan-liaofan)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.TPA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.TPA_DB.Create(&api).Error
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	err = global.TPA_DB.Delete(&api).Error
	CasbinServiceApp.ClearCasbin(1, api.Path, api.Method)
	return err
}

// GetAPIInfoList @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: GetAPIInfoList
// @description: 分页获取数据,
// @param: api model.SysApi, info request.PageInfo, order string, desc bool
// @return: err error
func (apiService *ApiService) GetAPIInfoList(api system.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.TPA_DB.Model(&system.SysApi{})
	var apiList []system.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi

func (apiService *ApiService) GetApiById(id uint) (err error, api system.SysApi) {
	err = global.TPA_DB.Where("id = ?", id).First(&api).Error
	return
}

// UpdateApi @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: UpdateApi
// @description: 根据id更新api
// @param: api model.SysApi
// @return: err error
func (apiService *ApiService) UpdateApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.TPA_DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.TPA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.TPA_DB.Save(&api).Error
		}
	}
	return err
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi

func (apiService *ApiService) GetAllApis() (err error, apis []system.SysApi) {
	err = global.TPA_DB.Find(&apis).Error
	return
}

// DeleteApisByIds @author: [liaofan](https://github.com/liaofan-liaofan)
// @function: DeleteApis
// @description: 删除选中API
// @param: apis []model.SysApi
// @return: err error
func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	err = global.TPA_DB.Delete(&[]system.SysApi{}, "id in ?", ids.Ids).Error
	return err
}
