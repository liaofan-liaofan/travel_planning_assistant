package source

import (
	"project/global"
	"project/model/system"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var menus = []system.SysBaseMenu{
	{TPA_MODEL: global.TPA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: system.Meta{Title: "仪表盘", Icon: "odometer"}},
	{TPA_MODEL: global.TPA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: system.Meta{Title: "关于我们", Icon: "info-filled"}},
	{TPA_MODEL: global.TPA_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: system.Meta{Title: "超级管理员", Icon: "user"}},
	{TPA_MODEL: global.TPA_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: system.Meta{Title: "角色管理", Icon: "avatar"}},
	{TPA_MODEL: global.TPA_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: system.Meta{Title: "菜单管理", Icon: "tickets", KeepAlive: true}},
	{TPA_MODEL: global.TPA_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: system.Meta{Title: "api管理", Icon: "platform", KeepAlive: true}},
	{TPA_MODEL: global.TPA_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: system.Meta{Title: "用户管理", Icon: "coordinate"}},
	{TPA_MODEL: global.TPA_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: system.Meta{Title: "个人信息", Icon: "message"}},
	{TPA_MODEL: global.TPA_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: system.Meta{Title: "服务器状态", Icon: "cloudy"}},
	{TPA_MODEL: global.TPA_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: system.Meta{Title: "操作历史", Icon: "pie-chart"}},
	{TPA_MODEL: global.TPA_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dept", Name: "dept", Component: "view/superAdmin/dept/dept.vue", Sort: 5, Meta: system.Meta{Title: "部门管理", Icon: "house"}},
}

// Init @author: [liaofan](https://github.com/liaofan-liaofan)
// @description: sys_base_menus 表数据初始化
func (m *menu) Init() error {
	return global.TPA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]system.SysBaseMenu{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_base_menus 表初始数据成功!")
		return nil
	})
}
