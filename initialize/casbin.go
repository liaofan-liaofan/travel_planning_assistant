package initialize

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	adapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"os"
	"project/global"
	"strings"
)

//@function: Casbin
//@description: 持久化到数据库  引入自定义规则
//@return: *casbin.Enforcer

func Casbin() *casbin.SyncedEnforcer {
	a, err := adapter.NewAdapterByDB(global.TPA_DB)
	syncedEnforcer, err := casbin.NewSyncedEnforcer(global.TPA_CONFIG.Casbin.ModelPath, a)
	syncedEnforcer.AddFunction("ParamsMatch", paramsMatchFunc)
	err = syncedEnforcer.LoadPolicy()
	if err != nil {
		global.TPA_LG.Error("Casbin 初始化异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	}
	return syncedEnforcer
}

// @function: ParamsMatchFunc
// @description: 自定义规则函数
// @param: args ...interface{}
// @return: interface{}, error
func paramsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return paramsMatch(name1, name2), nil
}

// @function: ParamsMatch
// @description: 自定义规则函数
// @param: fullNameKey1 string, key2 string
// @return: bool
func paramsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}
