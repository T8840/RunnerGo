// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameRole = "role"

// Role mapped from table <role>
type Role struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`      // 主键id
	RoleID    string         `gorm:"column:role_id;not null" json:"role_id"`                 // 角色id
	RoleType  int32          `gorm:"column:role_type;not null" json:"role_type"`             // 角色分类（1：企业  2：团队）
	Name      string         `gorm:"column:name;not null" json:"name"`                       // 角色名称
	CompanyID string         `gorm:"column:company_id;not null" json:"company_id"`           // 企业id
	Level     int32          `gorm:"column:level;not null" json:"level"`                     // 角色层级（1:超管/团队管理员 2:管理员/团队成员 3:普通成员/只读成员/自定义角色）
	IsDefault int32          `gorm:"column:is_default;not null;default:2" json:"is_default"` // 是否是默认角色  1：是   2：自定义角色
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
