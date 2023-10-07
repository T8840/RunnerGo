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

	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/model"
)

func newAutoPlanTaskConf(db *gorm.DB, opts ...gen.DOOption) autoPlanTaskConf {
	_autoPlanTaskConf := autoPlanTaskConf{}

	_autoPlanTaskConf.autoPlanTaskConfDo.UseDB(db, opts...)
	_autoPlanTaskConf.autoPlanTaskConfDo.UseModel(&model.AutoPlanTaskConf{})

	tableName := _autoPlanTaskConf.autoPlanTaskConfDo.TableName()
	_autoPlanTaskConf.ALL = field.NewAsterisk(tableName)
	_autoPlanTaskConf.ID = field.NewInt32(tableName, "id")
	_autoPlanTaskConf.PlanID = field.NewString(tableName, "plan_id")
	_autoPlanTaskConf.TeamID = field.NewString(tableName, "team_id")
	_autoPlanTaskConf.TaskType = field.NewInt32(tableName, "task_type")
	_autoPlanTaskConf.TaskMode = field.NewInt32(tableName, "task_mode")
	_autoPlanTaskConf.SceneRunOrder = field.NewInt32(tableName, "scene_run_order")
	_autoPlanTaskConf.TestCaseRunOrder = field.NewInt32(tableName, "test_case_run_order")
	_autoPlanTaskConf.RunUserID = field.NewString(tableName, "run_user_id")
	_autoPlanTaskConf.CreatedAt = field.NewTime(tableName, "created_at")
	_autoPlanTaskConf.UpdatedAt = field.NewTime(tableName, "updated_at")
	_autoPlanTaskConf.DeletedAt = field.NewField(tableName, "deleted_at")

	_autoPlanTaskConf.fillFieldMap()

	return _autoPlanTaskConf
}

type autoPlanTaskConf struct {
	autoPlanTaskConfDo autoPlanTaskConfDo

	ALL              field.Asterisk
	ID               field.Int32  // 配置ID
	PlanID           field.String // 计划ID
	TeamID           field.String // 团队ID
	TaskType         field.Int32  // 任务类型：1-普通模式，2-定时任务
	TaskMode         field.Int32  // 运行模式：1-按照用例执行
	SceneRunOrder    field.Int32  // 场景运行次序：1-顺序执行，2-同时执行
	TestCaseRunOrder field.Int32  // 用例运行次序：1-顺序执行，2-同时执行
	RunUserID        field.String // 运行人用户ID
	CreatedAt        field.Time   // 创建时间
	UpdatedAt        field.Time   // 更新时间
	DeletedAt        field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (a autoPlanTaskConf) Table(newTableName string) *autoPlanTaskConf {
	a.autoPlanTaskConfDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a autoPlanTaskConf) As(alias string) *autoPlanTaskConf {
	a.autoPlanTaskConfDo.DO = *(a.autoPlanTaskConfDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *autoPlanTaskConf) updateTableName(table string) *autoPlanTaskConf {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.PlanID = field.NewString(table, "plan_id")
	a.TeamID = field.NewString(table, "team_id")
	a.TaskType = field.NewInt32(table, "task_type")
	a.TaskMode = field.NewInt32(table, "task_mode")
	a.SceneRunOrder = field.NewInt32(table, "scene_run_order")
	a.TestCaseRunOrder = field.NewInt32(table, "test_case_run_order")
	a.RunUserID = field.NewString(table, "run_user_id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *autoPlanTaskConf) WithContext(ctx context.Context) *autoPlanTaskConfDo {
	return a.autoPlanTaskConfDo.WithContext(ctx)
}

func (a autoPlanTaskConf) TableName() string { return a.autoPlanTaskConfDo.TableName() }

func (a autoPlanTaskConf) Alias() string { return a.autoPlanTaskConfDo.Alias() }

func (a *autoPlanTaskConf) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *autoPlanTaskConf) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 11)
	a.fieldMap["id"] = a.ID
	a.fieldMap["plan_id"] = a.PlanID
	a.fieldMap["team_id"] = a.TeamID
	a.fieldMap["task_type"] = a.TaskType
	a.fieldMap["task_mode"] = a.TaskMode
	a.fieldMap["scene_run_order"] = a.SceneRunOrder
	a.fieldMap["test_case_run_order"] = a.TestCaseRunOrder
	a.fieldMap["run_user_id"] = a.RunUserID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
}

func (a autoPlanTaskConf) clone(db *gorm.DB) autoPlanTaskConf {
	a.autoPlanTaskConfDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a autoPlanTaskConf) replaceDB(db *gorm.DB) autoPlanTaskConf {
	a.autoPlanTaskConfDo.ReplaceDB(db)
	return a
}

type autoPlanTaskConfDo struct{ gen.DO }

func (a autoPlanTaskConfDo) Debug() *autoPlanTaskConfDo {
	return a.withDO(a.DO.Debug())
}

func (a autoPlanTaskConfDo) WithContext(ctx context.Context) *autoPlanTaskConfDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a autoPlanTaskConfDo) ReadDB() *autoPlanTaskConfDo {
	return a.Clauses(dbresolver.Read)
}

func (a autoPlanTaskConfDo) WriteDB() *autoPlanTaskConfDo {
	return a.Clauses(dbresolver.Write)
}

func (a autoPlanTaskConfDo) Session(config *gorm.Session) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Session(config))
}

func (a autoPlanTaskConfDo) Clauses(conds ...clause.Expression) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a autoPlanTaskConfDo) Returning(value interface{}, columns ...string) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a autoPlanTaskConfDo) Not(conds ...gen.Condition) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a autoPlanTaskConfDo) Or(conds ...gen.Condition) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a autoPlanTaskConfDo) Select(conds ...field.Expr) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a autoPlanTaskConfDo) Where(conds ...gen.Condition) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a autoPlanTaskConfDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *autoPlanTaskConfDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a autoPlanTaskConfDo) Order(conds ...field.Expr) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a autoPlanTaskConfDo) Distinct(cols ...field.Expr) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a autoPlanTaskConfDo) Omit(cols ...field.Expr) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a autoPlanTaskConfDo) Join(table schema.Tabler, on ...field.Expr) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a autoPlanTaskConfDo) LeftJoin(table schema.Tabler, on ...field.Expr) *autoPlanTaskConfDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a autoPlanTaskConfDo) RightJoin(table schema.Tabler, on ...field.Expr) *autoPlanTaskConfDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a autoPlanTaskConfDo) Group(cols ...field.Expr) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a autoPlanTaskConfDo) Having(conds ...gen.Condition) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a autoPlanTaskConfDo) Limit(limit int) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a autoPlanTaskConfDo) Offset(offset int) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a autoPlanTaskConfDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a autoPlanTaskConfDo) Unscoped() *autoPlanTaskConfDo {
	return a.withDO(a.DO.Unscoped())
}

func (a autoPlanTaskConfDo) Create(values ...*model.AutoPlanTaskConf) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a autoPlanTaskConfDo) CreateInBatches(values []*model.AutoPlanTaskConf, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a autoPlanTaskConfDo) Save(values ...*model.AutoPlanTaskConf) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a autoPlanTaskConfDo) First() (*model.AutoPlanTaskConf, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoPlanTaskConf), nil
	}
}

func (a autoPlanTaskConfDo) Take() (*model.AutoPlanTaskConf, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoPlanTaskConf), nil
	}
}

func (a autoPlanTaskConfDo) Last() (*model.AutoPlanTaskConf, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoPlanTaskConf), nil
	}
}

func (a autoPlanTaskConfDo) Find() ([]*model.AutoPlanTaskConf, error) {
	result, err := a.DO.Find()
	return result.([]*model.AutoPlanTaskConf), err
}

func (a autoPlanTaskConfDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AutoPlanTaskConf, err error) {
	buf := make([]*model.AutoPlanTaskConf, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a autoPlanTaskConfDo) FindInBatches(result *[]*model.AutoPlanTaskConf, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a autoPlanTaskConfDo) Attrs(attrs ...field.AssignExpr) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a autoPlanTaskConfDo) Assign(attrs ...field.AssignExpr) *autoPlanTaskConfDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a autoPlanTaskConfDo) Joins(fields ...field.RelationField) *autoPlanTaskConfDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a autoPlanTaskConfDo) Preload(fields ...field.RelationField) *autoPlanTaskConfDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a autoPlanTaskConfDo) FirstOrInit() (*model.AutoPlanTaskConf, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoPlanTaskConf), nil
	}
}

func (a autoPlanTaskConfDo) FirstOrCreate() (*model.AutoPlanTaskConf, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoPlanTaskConf), nil
	}
}

func (a autoPlanTaskConfDo) FindByPage(offset int, limit int) (result []*model.AutoPlanTaskConf, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a autoPlanTaskConfDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a autoPlanTaskConfDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a autoPlanTaskConfDo) Delete(models ...*model.AutoPlanTaskConf) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *autoPlanTaskConfDo) withDO(do gen.Dao) *autoPlanTaskConfDo {
	a.DO = *do.(*gen.DO)
	return a
}
