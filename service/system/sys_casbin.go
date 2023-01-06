package system

import (
	"errors"
	"project/global"
	"project/model/system"
	"project/model/system/request"
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

// UpdateCasbinApi @function: UpdateCasbinApi
// @description: API更新随动
// @param: oldPath string, newPath string, oldMethod string, newMethod string
// @return: error
func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.TPA_DB.Table("casbin_rule").Model(&system.CasbinModel{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// UpdateCasbin @function: UpdateCasbin
// @description: 更新casbin权限
// @param: authorityId string, casbinInfos []request.CasbinInfo
// @return: error
func (casbinService *CasbinService) UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	casbinService.ClearCasbin(0, authorityId)
	var rules [][]string
	for _, v := range casbinInfos {
		cm := system.CasbinModel{
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	success, _ := global.TPA_CASBIN.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

// ClearCasbin @function: ClearCasbin
// @description: 清除匹配的权限
// @param: v int, p ...string
// @return: bool
func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	success, _ := global.TPA_CASBIN.RemoveFilteredPolicy(v, p...)
	return success
}

// GetPolicyPathByAuthorityId @function: GetPolicyPathByAuthorityId
// @description: 获取权限列表
// @param: authorityId string
// @return: pathMaps []request.CasbinInfo
func (casbinService *CasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	list := global.TPA_CASBIN.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}
