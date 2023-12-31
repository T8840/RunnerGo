// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameThirdNoticeGroup = "third_notice_group"

// ThirdNoticeGroup mapped from table <third_notice_group>
type ThirdNoticeGroup struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键id
	GroupID   string         `gorm:"column:group_id;not null" json:"group_id"`          // 通知组id
	Name      string         `gorm:"column:name;not null" json:"name"`                  // 通知组名称
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ThirdNoticeGroup's table name
func (*ThirdNoticeGroup) TableName() string {
	return TableNameThirdNoticeGroup
}
