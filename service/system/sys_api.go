package system

import (
	"errors"
	"gorm.io/gorm"
	"project/global"
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
