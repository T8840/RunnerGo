// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameStressPlanTaskConf = "stress_plan_task_conf"

// StressPlanTaskConf mapped from table <stress_plan_task_conf>
type StressPlanTaskConf struct {
	ID                      int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                            // 配置ID
	PlanID                  string         `gorm:"column:plan_id;not null;default:0" json:"plan_id"`                             // 计划ID
	TeamID                  string         `gorm:"column:team_id;not null" json:"team_id"`                                       // 团队ID
	SceneID                 string         `gorm:"column:scene_id;not null" json:"scene_id"`                                     // 场景ID
	TaskType                int32          `gorm:"column:task_type;not null" json:"task_type"`                                   // 任务类型：1-普通模式，2-定时任务
	TaskMode                int32          `gorm:"column:task_mode;not null" json:"task_mode"`                                   // 压测模式：1-并发模式，2-阶梯模式，3-错误率模式，4-响应时间模式，5-每秒请求数模式，6-每秒事务数模式
	ControlMode             int32          `gorm:"column:control_mode;not null" json:"control_mode"`                             // 控制模式：0-集中模式，1-单独模式
	DebugMode               string         `gorm:"column:debug_mode;not null;default:stop" json:"debug_mode"`                    // debug模式：stop-关闭，all-开启全部日志，only_success-开启仅成功日志，only_error-开启仅错误日志
	ModeConf                string         `gorm:"column:mode_conf;not null" json:"mode_conf"`                                   // 压测模式配置详情
	IsOpenDistributed       int32          `gorm:"column:is_open_distributed;not null" json:"is_open_distributed"`               // 是否开启分布式调度：0-关闭，1-开启
	MachineDispatchModeConf string         `gorm:"column:machine_dispatch_mode_conf;not null" json:"machine_dispatch_mode_conf"` // 分布式压力机配置
	RunUserID               string         `gorm:"column:run_user_id;not null;default:0" json:"run_user_id"`                     // 运行人用户ID
	CreatedAt               time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`       // 创建时间
	UpdatedAt               time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`       // 更新时间
	DeletedAt               gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                          // 删除时间
}

// TableName StressPlanTaskConf's table name
func (*StressPlanTaskConf) TableName() string {
	return TableNameStressPlanTaskConf
}
