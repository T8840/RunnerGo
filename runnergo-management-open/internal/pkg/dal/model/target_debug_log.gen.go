// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTargetDebugLog = "target_debug_log"

// TargetDebugLog mapped from table <target_debug_log>
type TargetDebugLog struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // 主键ID
	TargetID   string         `gorm:"column:target_id;not null" json:"target_id"`                             // 目标唯一ID
	TargetType int32          `gorm:"column:target_type;not null" json:"target_type"`                         // 目标类型：1-api，2-scene
	TeamID     string         `gorm:"column:team_id;not null" json:"team_id"`                                 // 团队ID
	CreatedAt  time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                    // 删除时间
}

// TableName TargetDebugLog's table name
func (*TargetDebugLog) TableName() string {
	return TableNameTargetDebugLog
}