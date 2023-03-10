package system

import "project/global"

type SysBaseMenu struct {
	global.TPA_MODEL
	MenuLevel     uint                              `json:"-"`
	ParentId      string                            `json:"parentId" gorm:"not null;comment:父菜单ID"`     // 父菜单ID
	Path          string                            `json:"path" gorm:"not null;comment:路由path"`        // 路由path
	Name          string                            `json:"name" gorm:"not null;comment:路由name"`        // 路由name
	Hidden        bool                              `json:"hidden" gorm:"not null;comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component     string                            `json:"component" gorm:"not null;comment:对应前端文件路径"` // 对应前端文件路径
	Sort          int                               `json:"sort" gorm:"not null;comment:排序标记"`          // 排序标记
	Meta          `json:"meta" gorm:"comment:附加属性"` // 附加属性
	SysAuthoritys []SysAuthority                    `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu                     `json:"children" gorm:"-"`
	Parameters    []SysBaseMenuParameter            `json:"parameters"`
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"not null;comment:是否缓存"`           // 是否缓存
	DefaultMenu bool   `json:"defaultMenu" gorm:"not null;comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
	Title       string `json:"title" gorm:"not null;comment:菜单名"`                // 菜单名
	Icon        string `json:"icon" gorm:"not null;comment:菜单图标"`                // 菜单图标
	CloseTab    bool   `json:"closeTab" gorm:"not null;comment:自动关闭tab"`         // 自动关闭tab
}

type SysBaseMenuParameter struct {
	global.TPA_MODEL
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"not null;comment:地址栏携带参数为params还是query"` // 地址栏携带参数为params还是query
	Key           string `json:"key" gorm:"not null;comment:地址栏携带参数的key"`            // 地址栏携带参数的key
	Value         string `json:"value" gorm:"not null;comment:地址栏携带参数的值"`            // 地址栏携带参数的值
}
