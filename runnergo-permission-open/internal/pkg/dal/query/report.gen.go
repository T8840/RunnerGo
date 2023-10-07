// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"permission-open/internal/pkg/dal/model"
)

func newReport(db *gorm.DB, opts ...gen.DOOption) report {
	_report := report{}

	_report.reportDo.UseDB(db, opts...)
	_report.reportDo.UseModel(&model.Report{})

	tableName := _report.reportDo.TableName()
	_report.ALL = field.NewAsterisk(tableName)
	_report.ID = field.NewInt64(tableName, "id")
	_report.TeamID = field.NewInt64(tableName, "team_id")
	_report.Rank = field.NewInt64(tableName, "rank")
	_report.PlanID = field.NewInt64(tableName, "plan_id")
	_report.PlanName = field.NewString(tableName, "plan_name")
	_report.SceneID = field.NewInt64(tableName, "scene_id")
	_report.SceneName = field.NewString(tableName, "scene_name")
	_report.TaskType = field.NewInt32(tableName, "task_type")
	_report.TaskMode = field.NewInt32(tableName, "task_mode")
	_report.Status = field.NewInt32(tableName, "status")
	_report.RanAt = field.NewTime(tableName, "ran_at")
	_report.RunUserIdentify = field.NewString(tableName, "run_user_identify")
	_report.RunUserID = field.NewInt64(tableName, "run_user_id")
	_report.CreatedAt = field.NewTime(tableName, "created_at")
	_report.UpdatedAt = field.NewTime(tableName, "updated_at")
	_report.DeletedAt = field.NewField(tableName, "deleted_at")

	_report.fillFieldMap()

	return _report
}

type report struct {
	reportDo reportDo

	ALL             field.Asterisk
	ID              field.Int64
	TeamID          field.Int64  // 团队ID
	Rank            field.Int64  // 团队内份数
	PlanID          field.Int64  // 计划ID
	PlanName        field.String // 计划名称
	SceneID         field.Int64  // 场景ID
	SceneName       field.String // 场景名称
	TaskType        field.Int32  // 任务类型
	TaskMode        field.Int32  // 压测模式
	Status          field.Int32  // 报告状态1:进行中，2:已完成
	RanAt           field.Time   // 启动时间
	RunUserIdentify field.String
	RunUserID       field.Int64 // 启动人id
	CreatedAt       field.Time
	UpdatedAt       field.Time
	DeletedAt       field.Field

	fieldMap map[string]field.Expr
}

func (r report) Table(newTableName string) *report {
	r.reportDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r report) As(alias string) *report {
	r.reportDo.DO = *(r.reportDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *report) updateTableName(table string) *report {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.TeamID = field.NewInt64(table, "team_id")
	r.Rank = field.NewInt64(table, "rank")
	r.PlanID = field.NewInt64(table, "plan_id")
	r.PlanName = field.NewString(table, "plan_name")
	r.SceneID = field.NewInt64(table, "scene_id")
	r.SceneName = field.NewString(table, "scene_name")
	r.TaskType = field.NewInt32(table, "task_type")
	r.TaskMode = field.NewInt32(table, "task_mode")
	r.Status = field.NewInt32(table, "status")
	r.RanAt = field.NewTime(table, "ran_at")
	r.RunUserIdentify = field.NewString(table, "run_user_identify")
	r.RunUserID = field.NewInt64(table, "run_user_id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")

	r.fillFieldMap()

	return r
}

func (r *report) WithContext(ctx context.Context) *reportDo { return r.reportDo.WithContext(ctx) }

func (r report) TableName() string { return r.reportDo.TableName() }

func (r report) Alias() string { return r.reportDo.Alias() }

func (r *report) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *report) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 16)
	r.fieldMap["id"] = r.ID
	r.fieldMap["team_id"] = r.TeamID
	r.fieldMap["rank"] = r.Rank
	r.fieldMap["plan_id"] = r.PlanID
	r.fieldMap["plan_name"] = r.PlanName
	r.fieldMap["scene_id"] = r.SceneID
	r.fieldMap["scene_name"] = r.SceneName
	r.fieldMap["task_type"] = r.TaskType
	r.fieldMap["task_mode"] = r.TaskMode
	r.fieldMap["status"] = r.Status
	r.fieldMap["ran_at"] = r.RanAt
	r.fieldMap["run_user_identify"] = r.RunUserIdentify
	r.fieldMap["run_user_id"] = r.RunUserID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
}

func (r report) clone(db *gorm.DB) report {
	r.reportDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r report) replaceDB(db *gorm.DB) report {
	r.reportDo.ReplaceDB(db)
	return r
}

type reportDo struct{ gen.DO }

func (r reportDo) Debug() *reportDo {
	return r.withDO(r.DO.Debug())
}

func (r reportDo) WithContext(ctx context.Context) *reportDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r reportDo) ReadDB() *reportDo {
	return r.Clauses(dbresolver.Read)
}

func (r reportDo) WriteDB() *reportDo {
	return r.Clauses(dbresolver.Write)
}

func (r reportDo) Session(config *gorm.Session) *reportDo {
	return r.withDO(r.DO.Session(config))
}

func (r reportDo) Clauses(conds ...clause.Expression) *reportDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r reportDo) Returning(value interface{}, columns ...string) *reportDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r reportDo) Not(conds ...gen.Condition) *reportDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r reportDo) Or(conds ...gen.Condition) *reportDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r reportDo) Select(conds ...field.Expr) *reportDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r reportDo) Where(conds ...gen.Condition) *reportDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r reportDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *reportDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r reportDo) Order(conds ...field.Expr) *reportDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r reportDo) Distinct(cols ...field.Expr) *reportDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r reportDo) Omit(cols ...field.Expr) *reportDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r reportDo) Join(table schema.Tabler, on ...field.Expr) *reportDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r reportDo) LeftJoin(table schema.Tabler, on ...field.Expr) *reportDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r reportDo) RightJoin(table schema.Tabler, on ...field.Expr) *reportDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r reportDo) Group(cols ...field.Expr) *reportDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r reportDo) Having(conds ...gen.Condition) *reportDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r reportDo) Limit(limit int) *reportDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r reportDo) Offset(offset int) *reportDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r reportDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *reportDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r reportDo) Unscoped() *reportDo {
	return r.withDO(r.DO.Unscoped())
}

func (r reportDo) Create(values ...*model.Report) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r reportDo) CreateInBatches(values []*model.Report, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r reportDo) Save(values ...*model.Report) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r reportDo) First() (*model.Report, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Report), nil
	}
}

func (r reportDo) Take() (*model.Report, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Report), nil
	}
}

func (r reportDo) Last() (*model.Report, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Report), nil
	}
}

func (r reportDo) Find() ([]*model.Report, error) {
	result, err := r.DO.Find()
	return result.([]*model.Report), err
}

func (r reportDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Report, err error) {
	buf := make([]*model.Report, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r reportDo) FindInBatches(result *[]*model.Report, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r reportDo) Attrs(attrs ...field.AssignExpr) *reportDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r reportDo) Assign(attrs ...field.AssignExpr) *reportDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r reportDo) Joins(fields ...field.RelationField) *reportDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r reportDo) Preload(fields ...field.RelationField) *reportDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r reportDo) FirstOrInit() (*model.Report, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Report), nil
	}
}

func (r reportDo) FirstOrCreate() (*model.Report, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Report), nil
	}
}

func (r reportDo) FindByPage(offset int, limit int) (result []*model.Report, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r reportDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r reportDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r reportDo) Delete(models ...*model.Report) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *reportDo) withDO(do gen.Dao) *reportDo {
	r.DO = *do.(*gen.DO)
	return r
}
