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

func newVariableImport(db *gorm.DB, opts ...gen.DOOption) variableImport {
	_variableImport := variableImport{}

	_variableImport.variableImportDo.UseDB(db, opts...)
	_variableImport.variableImportDo.UseModel(&model.VariableImport{})

	tableName := _variableImport.variableImportDo.TableName()
	_variableImport.ALL = field.NewAsterisk(tableName)
	_variableImport.ID = field.NewInt64(tableName, "id")
	_variableImport.TeamID = field.NewString(tableName, "team_id")
	_variableImport.SceneID = field.NewString(tableName, "scene_id")
	_variableImport.Name = field.NewString(tableName, "name")
	_variableImport.URL = field.NewString(tableName, "url")
	_variableImport.UploaderID = field.NewString(tableName, "uploader_id")
	_variableImport.Status = field.NewInt32(tableName, "status")
	_variableImport.CreatedAt = field.NewTime(tableName, "created_at")
	_variableImport.UpdatedAt = field.NewTime(tableName, "updated_at")
	_variableImport.DeletedAt = field.NewField(tableName, "deleted_at")

	_variableImport.fillFieldMap()

	return _variableImport
}

type variableImport struct {
	variableImportDo variableImportDo

	ALL        field.Asterisk
	ID         field.Int64
	TeamID     field.String // 团队id
	SceneID    field.String // 场景id
	Name       field.String // 文件名称
	URL        field.String // 文件地址
	UploaderID field.String // 上传人id
	Status     field.Int32  // 开关状态：1-开，2-关
	CreatedAt  field.Time   // 创建时间
	UpdatedAt  field.Time   // 修改时间
	DeletedAt  field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (v variableImport) Table(newTableName string) *variableImport {
	v.variableImportDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v variableImport) As(alias string) *variableImport {
	v.variableImportDo.DO = *(v.variableImportDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *variableImport) updateTableName(table string) *variableImport {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewInt64(table, "id")
	v.TeamID = field.NewString(table, "team_id")
	v.SceneID = field.NewString(table, "scene_id")
	v.Name = field.NewString(table, "name")
	v.URL = field.NewString(table, "url")
	v.UploaderID = field.NewString(table, "uploader_id")
	v.Status = field.NewInt32(table, "status")
	v.CreatedAt = field.NewTime(table, "created_at")
	v.UpdatedAt = field.NewTime(table, "updated_at")
	v.DeletedAt = field.NewField(table, "deleted_at")

	v.fillFieldMap()

	return v
}

func (v *variableImport) WithContext(ctx context.Context) *variableImportDo {
	return v.variableImportDo.WithContext(ctx)
}

func (v variableImport) TableName() string { return v.variableImportDo.TableName() }

func (v variableImport) Alias() string { return v.variableImportDo.Alias() }

func (v *variableImport) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *variableImport) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 10)
	v.fieldMap["id"] = v.ID
	v.fieldMap["team_id"] = v.TeamID
	v.fieldMap["scene_id"] = v.SceneID
	v.fieldMap["name"] = v.Name
	v.fieldMap["url"] = v.URL
	v.fieldMap["uploader_id"] = v.UploaderID
	v.fieldMap["status"] = v.Status
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["updated_at"] = v.UpdatedAt
	v.fieldMap["deleted_at"] = v.DeletedAt
}

func (v variableImport) clone(db *gorm.DB) variableImport {
	v.variableImportDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v variableImport) replaceDB(db *gorm.DB) variableImport {
	v.variableImportDo.ReplaceDB(db)
	return v
}

type variableImportDo struct{ gen.DO }

func (v variableImportDo) Debug() *variableImportDo {
	return v.withDO(v.DO.Debug())
}

func (v variableImportDo) WithContext(ctx context.Context) *variableImportDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v variableImportDo) ReadDB() *variableImportDo {
	return v.Clauses(dbresolver.Read)
}

func (v variableImportDo) WriteDB() *variableImportDo {
	return v.Clauses(dbresolver.Write)
}

func (v variableImportDo) Session(config *gorm.Session) *variableImportDo {
	return v.withDO(v.DO.Session(config))
}

func (v variableImportDo) Clauses(conds ...clause.Expression) *variableImportDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v variableImportDo) Returning(value interface{}, columns ...string) *variableImportDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v variableImportDo) Not(conds ...gen.Condition) *variableImportDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v variableImportDo) Or(conds ...gen.Condition) *variableImportDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v variableImportDo) Select(conds ...field.Expr) *variableImportDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v variableImportDo) Where(conds ...gen.Condition) *variableImportDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v variableImportDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *variableImportDo {
	return v.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (v variableImportDo) Order(conds ...field.Expr) *variableImportDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v variableImportDo) Distinct(cols ...field.Expr) *variableImportDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v variableImportDo) Omit(cols ...field.Expr) *variableImportDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v variableImportDo) Join(table schema.Tabler, on ...field.Expr) *variableImportDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v variableImportDo) LeftJoin(table schema.Tabler, on ...field.Expr) *variableImportDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v variableImportDo) RightJoin(table schema.Tabler, on ...field.Expr) *variableImportDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v variableImportDo) Group(cols ...field.Expr) *variableImportDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v variableImportDo) Having(conds ...gen.Condition) *variableImportDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v variableImportDo) Limit(limit int) *variableImportDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v variableImportDo) Offset(offset int) *variableImportDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v variableImportDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *variableImportDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v variableImportDo) Unscoped() *variableImportDo {
	return v.withDO(v.DO.Unscoped())
}

func (v variableImportDo) Create(values ...*model.VariableImport) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v variableImportDo) CreateInBatches(values []*model.VariableImport, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v variableImportDo) Save(values ...*model.VariableImport) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v variableImportDo) First() (*model.VariableImport, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VariableImport), nil
	}
}

func (v variableImportDo) Take() (*model.VariableImport, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VariableImport), nil
	}
}

func (v variableImportDo) Last() (*model.VariableImport, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VariableImport), nil
	}
}

func (v variableImportDo) Find() ([]*model.VariableImport, error) {
	result, err := v.DO.Find()
	return result.([]*model.VariableImport), err
}

func (v variableImportDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VariableImport, err error) {
	buf := make([]*model.VariableImport, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v variableImportDo) FindInBatches(result *[]*model.VariableImport, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v variableImportDo) Attrs(attrs ...field.AssignExpr) *variableImportDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v variableImportDo) Assign(attrs ...field.AssignExpr) *variableImportDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v variableImportDo) Joins(fields ...field.RelationField) *variableImportDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v variableImportDo) Preload(fields ...field.RelationField) *variableImportDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v variableImportDo) FirstOrInit() (*model.VariableImport, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VariableImport), nil
	}
}

func (v variableImportDo) FirstOrCreate() (*model.VariableImport, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VariableImport), nil
	}
}

func (v variableImportDo) FindByPage(offset int, limit int) (result []*model.VariableImport, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v variableImportDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v variableImportDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v variableImportDo) Delete(models ...*model.VariableImport) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *variableImportDo) withDO(do gen.Dao) *variableImportDo {
	v.DO = *do.(*gen.DO)
	return v
}
