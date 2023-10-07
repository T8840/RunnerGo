// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAutoPlanTimedTaskConf = "auto_plan_timed_task_conf"

// AutoPlanTimedTaskConf mapped from table <auto_plan_timed_task_conf>
type AutoPlanTimedTaskConf struct {
	ID               int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 表id
	PlanID           string         `gorm:"column:plan_id;not null;default:0" json:"plan_id"`                         // 计划id
	TeamID           string         `gorm:"column:team_id;not null" json:"team_id"`                                   // 团队id
	Frequency        int32          `gorm:"column:frequency;not null" json:"frequency"`                               // 任务执行频次: 0-一次，1-每天，2-每周，3-每月
	TaskExecTime     int64          `gorm:"column:task_exec_time;not null" json:"task_exec_time"`                     // 任务执行时间
	TaskCloseTime    int64          `gorm:"column:task_close_time;not null" json:"task_close_time"`                   // 任务结束时间
	TaskType         int32          `gorm:"column:task_type;not null;default:2" json:"task_type"`                     // 任务类型：1-普通任务，2-定时任务
	TaskMode         int32          `gorm:"column:task_mode;not null;default:1" json:"task_mode"`                     // 运行模式：1-按照用例执行
	SceneRunOrder    int32          `gorm:"column:scene_run_order;not null;default:1" json:"scene_run_order"`         // 场景运行次序：1-顺序执行，2-同时执行
	TestCaseRunOrder int32          `gorm:"column:test_case_run_order;not null;default:1" json:"test_case_run_order"` // 测试用例运行次序：1-顺序执行，2-同时执行
	Status           int32          `gorm:"column:status;not null" json:"status"`                                     // 任务状态：0-未启用，1-运行中，2-已过期
	RunUserID        string         `gorm:"column:run_user_id;not null;default:0" json:"run_user_id"`                 // 运行人用户ID
	CreatedAt        time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`   // 创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`   // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                      // 删除时间
}

// TableName AutoPlanTimedTaskConf's table name
func (*AutoPlanTimedTaskConf) TableName() string {
	return TableNameAutoPlanTimedTaskConf
}
