package global

import (
	"gorm.io/gorm"
	"time"
)

type TPA_MODEL struct {
	ID        uint           `gorm:"primarykey;autoIncrement"` // 主键ID
	CreateBy  uint           //创建人
	UpdateBy  uint           //更新人
	CreatedAt time.Time      `json:"created_at"`      // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"-"` // 删除时间
}
