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

func newTeamEnvMysql(db *gorm.DB, opts ...gen.DOOption) teamEnvMysql {
	_teamEnvMysql := teamEnvMysql{}

	_teamEnvMysql.teamEnvMysqlDo.UseDB(db, opts...)
	_teamEnvMysql.teamEnvMysqlDo.UseModel(&model.TeamEnvMysql{})

	tableName := _teamEnvMysql.teamEnvMysqlDo.TableName()
	_teamEnvMysql.ALL = field.NewAsterisk(tableName)
	_teamEnvMysql.ID = field.NewInt64(tableName, "id")
	_teamEnvMysql.TeamID = field.NewString(tableName, "team_id")
	_teamEnvMysql.TeamEnvID = field.NewInt64(tableName, "team_env_id")
	_teamEnvMysql.Type = field.NewInt32(tableName, "type")
	_teamEnvMysql.ServerName = field.NewString(tableName, "server_name")
	_teamEnvMysql.Host = field.NewString(tableName, "host")
	_teamEnvMysql.Port = field.NewInt32(tableName, "port")
	_teamEnvMysql.User = field.NewString(tableName, "user")
	_teamEnvMysql.Password = field.NewString(tableName, "password")
	_teamEnvMysql.DbName = field.NewString(tableName, "db_name")
	_teamEnvMysql.Charset = field.NewString(tableName, "charset")
	_teamEnvMysql.CreatedAt = field.NewTime(tableName, "created_at")
	_teamEnvMysql.UpdatedAt = field.NewTime(tableName, "updated_at")
	_teamEnvMysql.DeletedAt = field.NewField(tableName, "deleted_at")

	_teamEnvMysql.fillFieldMap()

	return _teamEnvMysql
}

type teamEnvMysql struct {
	teamEnvMysqlDo teamEnvMysqlDo

	ALL        field.Asterisk
	ID         field.Int64  // 主键id
	TeamID     field.String // 团队id
	TeamEnvID  field.Int64  // 环境变量id
	Type       field.Int32  // 数据库类型：1-mysql，2-sql server
	ServerName field.String // mysql服务名称
	Host       field.String // 服务地址
	Port       field.Int32  // 端口号
	User       field.String // 账号
	Password   field.String // 密码
	DbName     field.String // 数据库名称
	Charset    field.String // 字符编码集
	CreatedAt  field.Time   // 创建时间
	UpdatedAt  field.Time   // 更新时间
	DeletedAt  field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (t teamEnvMysql) Table(newTableName string) *teamEnvMysql {
	t.teamEnvMysqlDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t teamEnvMysql) As(alias string) *teamEnvMysql {
	t.teamEnvMysqlDo.DO = *(t.teamEnvMysqlDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *teamEnvMysql) updateTableName(table string) *teamEnvMysql {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.TeamID = field.NewString(table, "team_id")
	t.TeamEnvID = field.NewInt64(table, "team_env_id")
	t.Type = field.NewInt32(table, "type")
	t.ServerName = field.NewString(table, "server_name")
	t.Host = field.NewString(table, "host")
	t.Port = field.NewInt32(table, "port")
	t.User = field.NewString(table, "user")
	t.Password = field.NewString(table, "password")
	t.DbName = field.NewString(table, "db_name")
	t.Charset = field.NewString(table, "charset")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *teamEnvMysql) WithContext(ctx context.Context) *teamEnvMysqlDo {
	return t.teamEnvMysqlDo.WithContext(ctx)
}

func (t teamEnvMysql) TableName() string { return t.teamEnvMysqlDo.TableName() }

func (t teamEnvMysql) Alias() string { return t.teamEnvMysqlDo.Alias() }

func (t *teamEnvMysql) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *teamEnvMysql) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 14)
	t.fieldMap["id"] = t.ID
	t.fieldMap["team_id"] = t.TeamID
	t.fieldMap["team_env_id"] = t.TeamEnvID
	t.fieldMap["type"] = t.Type
	t.fieldMap["server_name"] = t.ServerName
	t.fieldMap["host"] = t.Host
	t.fieldMap["port"] = t.Port
	t.fieldMap["user"] = t.User
	t.fieldMap["password"] = t.Password
	t.fieldMap["db_name"] = t.DbName
	t.fieldMap["charset"] = t.Charset
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t teamEnvMysql) clone(db *gorm.DB) teamEnvMysql {
	t.teamEnvMysqlDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t teamEnvMysql) replaceDB(db *gorm.DB) teamEnvMysql {
	t.teamEnvMysqlDo.ReplaceDB(db)
	return t
}

type teamEnvMysqlDo struct{ gen.DO }

func (t teamEnvMysqlDo) Debug() *teamEnvMysqlDo {
	return t.withDO(t.DO.Debug())
}

func (t teamEnvMysqlDo) WithContext(ctx context.Context) *teamEnvMysqlDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t teamEnvMysqlDo) ReadDB() *teamEnvMysqlDo {
	return t.Clauses(dbresolver.Read)
}

func (t teamEnvMysqlDo) WriteDB() *teamEnvMysqlDo {
	return t.Clauses(dbresolver.Write)
}

func (t teamEnvMysqlDo) Session(config *gorm.Session) *teamEnvMysqlDo {
	return t.withDO(t.DO.Session(config))
}

func (t teamEnvMysqlDo) Clauses(conds ...clause.Expression) *teamEnvMysqlDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t teamEnvMysqlDo) Returning(value interface{}, columns ...string) *teamEnvMysqlDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t teamEnvMysqlDo) Not(conds ...gen.Condition) *teamEnvMysqlDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t teamEnvMysqlDo) Or(conds ...gen.Condition) *teamEnvMysqlDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t teamEnvMysqlDo) Select(conds ...field.Expr) *teamEnvMysqlDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t teamEnvMysqlDo) Where(conds ...gen.Condition) *teamEnvMysqlDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t teamEnvMysqlDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *teamEnvMysqlDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t teamEnvMysqlDo) Order(conds ...field.Expr) *teamEnvMysqlDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t teamEnvMysqlDo) Distinct(cols ...field.Expr) *teamEnvMysqlDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t teamEnvMysqlDo) Omit(cols ...field.Expr) *teamEnvMysqlDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t teamEnvMysqlDo) Join(table schema.Tabler, on ...field.Expr) *teamEnvMysqlDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t teamEnvMysqlDo) LeftJoin(table schema.Tabler, on ...field.Expr) *teamEnvMysqlDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t teamEnvMysqlDo) RightJoin(table schema.Tabler, on ...field.Expr) *teamEnvMysqlDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t teamEnvMysqlDo) Group(cols ...field.Expr) *teamEnvMysqlDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t teamEnvMysqlDo) Having(conds ...gen.Condition) *teamEnvMysqlDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t teamEnvMysqlDo) Limit(limit int) *teamEnvMysqlDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t teamEnvMysqlDo) Offset(offset int) *teamEnvMysqlDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t teamEnvMysqlDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *teamEnvMysqlDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t teamEnvMysqlDo) Unscoped() *teamEnvMysqlDo {
	return t.withDO(t.DO.Unscoped())
}

func (t teamEnvMysqlDo) Create(values ...*model.TeamEnvMysql) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t teamEnvMysqlDo) CreateInBatches(values []*model.TeamEnvMysql, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t teamEnvMysqlDo) Save(values ...*model.TeamEnvMysql) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t teamEnvMysqlDo) First() (*model.TeamEnvMysql, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamEnvMysql), nil
	}
}

func (t teamEnvMysqlDo) Take() (*model.TeamEnvMysql, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamEnvMysql), nil
	}
}

func (t teamEnvMysqlDo) Last() (*model.TeamEnvMysql, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamEnvMysql), nil
	}
}

func (t teamEnvMysqlDo) Find() ([]*model.TeamEnvMysql, error) {
	result, err := t.DO.Find()
	return result.([]*model.TeamEnvMysql), err
}

func (t teamEnvMysqlDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TeamEnvMysql, err error) {
	buf := make([]*model.TeamEnvMysql, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t teamEnvMysqlDo) FindInBatches(result *[]*model.TeamEnvMysql, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t teamEnvMysqlDo) Attrs(attrs ...field.AssignExpr) *teamEnvMysqlDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t teamEnvMysqlDo) Assign(attrs ...field.AssignExpr) *teamEnvMysqlDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t teamEnvMysqlDo) Joins(fields ...field.RelationField) *teamEnvMysqlDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t teamEnvMysqlDo) Preload(fields ...field.RelationField) *teamEnvMysqlDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t teamEnvMysqlDo) FirstOrInit() (*model.TeamEnvMysql, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamEnvMysql), nil
	}
}

func (t teamEnvMysqlDo) FirstOrCreate() (*model.TeamEnvMysql, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamEnvMysql), nil
	}
}

func (t teamEnvMysqlDo) FindByPage(offset int, limit int) (result []*model.TeamEnvMysql, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t teamEnvMysqlDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t teamEnvMysqlDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t teamEnvMysqlDo) Delete(models ...*model.TeamEnvMysql) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *teamEnvMysqlDo) withDO(do gen.Dao) *teamEnvMysqlDo {
	t.DO = *do.(*gen.DO)
	return t
}
