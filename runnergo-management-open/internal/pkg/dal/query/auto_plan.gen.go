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

	"github.com/Runner-Go-Team/RunnerGo-management-open/internal/pkg/dal/model"
)

func newAutoPlan(db *gorm.DB, opts ...gen.DOOption) autoPlan {
	_autoPlan := autoPlan{}

	_autoPlan.autoPlanDo.UseDB(db, opts...)
	_autoPlan.autoPlanDo.UseModel(&model.AutoPlan{})

	tableName := _autoPlan.autoPlanDo.TableName()
	_autoPlan.ALL = field.NewAsterisk(tableName)
	_autoPlan.ID = field.NewInt64(tableName, "id")
	_autoPlan.PlanID = field.NewString(tableName, "plan_id")
	_autoPlan.RankID = field.NewInt64(tableName, "rank_id")
	_autoPlan.TeamID = field.NewString(tableName, "team_id")
	_autoPlan.PlanName = field.NewString(tableName, "plan_name")
	_autoPlan.TaskType = field.NewInt32(tableName, "task_type")
	_autoPlan.Status = field.NewInt32(tableName, "status")
	_autoPlan.CreateUserID = field.NewString(tableName, "create_user_id")
	_autoPlan.RunUserID = field.NewString(tableName, "run_user_id")
	_autoPlan.Remark = field.NewString(tableName, "remark")
	_autoPlan.RunCount = field.NewInt64(tableName, "run_count")
	_autoPlan.CreatedAt = field.NewTime(tableName, "created_at")
	_autoPlan.UpdatedAt = field.NewTime(tableName, "updated_at")
	_autoPlan.DeletedAt = field.NewField(tableName, "deleted_at")

	_autoPlan.fillFieldMap()

	return _autoPlan
}

type autoPlan struct {
	autoPlanDo autoPlanDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键
	PlanID       field.String // 计划ID
	RankID       field.Int64  // 序号ID
	TeamID       field.String // 团队ID
	PlanName     field.String // 计划名称
	TaskType     field.Int32  // 计划类型：1-普通任务，2-定时任务
	Status       field.Int32  // 计划状：1-未开始，2-进行中
	CreateUserID field.String // 创建人id
	RunUserID    field.String // 运行人id
	Remark       field.String // 备注
	RunCount     field.Int64  // 运行次数
	CreatedAt    field.Time   // 创建时间
	UpdatedAt    field.Time   // 更新时间
	DeletedAt    field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (a autoPlan) Table(newTableName string) *autoPlan {
	a.autoPlanDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a autoPlan) As(alias string) *autoPlan {
	a.autoPlanDo.DO = *(a.autoPlanDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *autoPlan) updateTableName(table string) *autoPlan {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.PlanID = field.NewString(table, "plan_id")
	a.RankID = field.NewInt64(table, "rank_id")
	a.TeamID = field.NewString(table, "team_id")
	a.PlanName = field.NewString(table, "plan_name")
	a.TaskType = field.NewInt32(table, "task_type")
	a.Status = field.NewInt32(table, "status")
	a.CreateUserID = field.NewString(table, "create_user_id")
	a.RunUserID = field.NewString(table, "run_user_id")
	a.Remark = field.NewString(table, "remark")
	a.RunCount = field.NewInt64(table, "run_count")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *autoPlan) WithContext(ctx context.Context) *autoPlanDo { return a.autoPlanDo.WithContext(ctx) }

func (a autoPlan) TableName() string { return a.autoPlanDo.TableName() }

func (a autoPlan) Alias() string { return a.autoPlanDo.Alias() }

func (a *autoPlan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *autoPlan) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 14)
	a.fieldMap["id"] = a.ID
	a.fieldMap["plan_id"] = a.PlanID
	a.fieldMap["rank_id"] = a.RankID
	a.fieldMap["team_id"] = a.TeamID
	a.fieldMap["plan_name"] = a.PlanName
	a.fieldMap["task_type"] = a.TaskType
	a.fieldMap["status"] = a.Status
	a.fieldMap["create_user_id"] = a.CreateUserID
	a.fieldMap["run_user_id"] = a.RunUserID
	a.fieldMap["remark"] = a.Remark
	a.fieldMap["run_count"] = a.RunCount
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
}

func (a autoPlan) clone(db *gorm.DB) autoPlan {
	a.autoPlanDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a autoPlan) replaceDB(db *gorm.DB) autoPlan {
	a.autoPlanDo.ReplaceDB(db)
	return a
}

type autoPlanDo struct{ gen.DO }

func (a autoPlanDo) Debug() *autoPlanDo {
	return a.withDO(a.DO.Debug())
}

func (a autoPlanDo) WithContext(ctx context.Context) *autoPlanDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a autoPlanDo) ReadDB() *autoPlanDo {
	return a.Clauses(dbresolver.Read)
}

func (a autoPlanDo) WriteDB() *autoPlanDo {
	return a.Clauses(dbresolver.Write)
}

func (a autoPlanDo) Session(config *gorm.Session) *autoPlanDo {
	return a.withDO(a.DO.Session(config))
}

func (a autoPlanDo) Clauses(conds ...clause.Expression) *autoPlanDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a autoPlanDo) Returning(value interface{}, columns ...string) *autoPlanDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a autoPlanDo) Not(conds ...gen.Condition) *autoPlanDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a autoPlanDo) Or(conds ...gen.Condition) *autoPlanDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a autoPlanDo) Select(conds ...field.Expr) *autoPlanDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a autoPlanDo) Where(conds ...gen.Condition) *autoPlanDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a autoPlanDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *autoPlanDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a autoPlanDo) Order(conds ...field.Expr) *autoPlanDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a autoPlanDo) Distinct(cols ...field.Expr) *autoPlanDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a autoPlanDo) Omit(cols ...field.Expr) *autoPlanDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a autoPlanDo) Join(table schema.Tabler, on ...field.Expr) *autoPlanDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a autoPlanDo) LeftJoin(table schema.Tabler, on ...field.Expr) *autoPlanDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a autoPlanDo) RightJoin(table schema.Tabler, on ...field.Expr) *autoPlanDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a autoPlanDo) Group(cols ...field.Expr) *autoPlanDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a autoPlanDo) Having(conds ...gen.Condition) *autoPlanDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a autoPlanDo) Limit(limit int) *autoPlanDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a autoPlanDo) Offset(offset int) *autoPlanDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a autoPlanDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *autoPlanDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a autoPlanDo) Unscoped() *autoPlanDo {
	return a.withDO(a.DO.Unscoped())
}

func (a autoPlanDo) Create(values ...*model.AutoPlan) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a autoPlanDo) CreateInBatches(values []*model.AutoPlan, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a autoPlanDo) Save(values ...*model.AutoPlan) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a autoPlanDo) First() (*model.AutoPlan, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoPlan), nil
	}
}

func (a autoPlanDo) Take() (*model.AutoPlan, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoPlan), nil
	}
}

func (a autoPlanDo) Last() (*model.AutoPlan, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoPlan), nil
	}
}

func (a autoPlanDo) Find() ([]*model.AutoPlan, error) {
	result, err := a.DO.Find()
	return result.([]*model.AutoPlan), err
}

func (a autoPlanDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AutoPlan, err error) {
	buf := make([]*model.AutoPlan, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a autoPlanDo) FindInBatches(result *[]*model.AutoPlan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a autoPlanDo) Attrs(attrs ...field.AssignExpr) *autoPlanDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a autoPlanDo) Assign(attrs ...field.AssignExpr) *autoPlanDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a autoPlanDo) Joins(fields ...field.RelationField) *autoPlanDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a autoPlanDo) Preload(fields ...field.RelationField) *autoPlanDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a autoPlanDo) FirstOrInit() (*model.AutoPlan, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoPlan), nil
	}
}

func (a autoPlanDo) FirstOrCreate() (*model.AutoPlan, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoPlan), nil
	}
}

func (a autoPlanDo) FindByPage(offset int, limit int) (result []*model.AutoPlan, count int64, err error) {
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

func (a autoPlanDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a autoPlanDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a autoPlanDo) Delete(models ...*model.AutoPlan) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *autoPlanDo) withDO(do gen.Dao) *autoPlanDo {
	a.DO = *do.(*gen.DO)
	return a
}
