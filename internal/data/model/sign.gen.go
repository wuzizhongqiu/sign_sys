// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSign = "sign"

// Sign 签到表
type Sign struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                          // 主键
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time  `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	DeleteAt   *time.Time `gorm:"column:delete_at;comment:逻辑删除标记" json:"delete_at"`                                      // 逻辑删除标记
	AttendName string     `gorm:"column:attend_name;not null;comment:签到" json:"attend_name"`                             // 签到
	SignTime   int32      `gorm:"column:sign_time;not null;comment:签到限时" json:"sign_time"`                               // 签到限时
}

// TableName Sign's table name
func (*Sign) TableName() string {
	return TableNameSign
}
