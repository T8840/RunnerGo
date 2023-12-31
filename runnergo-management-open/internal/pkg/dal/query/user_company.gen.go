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

func newUserCompany(db *gorm.DB, opts ...gen.DOOption) userCompany {
	_userCompany := userCompany{}

	_userCompany.userCompanyDo.UseDB(db, opts...)
	_userCompany.userCompanyDo.UseModel(&model.UserCompany{})

	tableName := _userCompany.userCompanyDo.TableName()
	_userCompany.ALL = field.NewAsterisk(tableName)
	_userCompany.ID = field.NewInt64(tableName, "id")
	_userCompany.UserID = field.NewString(tableName, "user_id")
	_userCompany.CompanyID = field.NewString(tableName, "company_id")
	_userCompany.InviteUserID = field.NewString(tableName, "invite_user_id")
	_userCompany.InviteTime = field.NewTime(tableName, "invite_time")
	_userCompany.Status = field.NewInt32(tableName, "status")
	_userCompany.CreatedAt = field.NewTime(tableName, "created_at")
	_userCompany.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userCompany.DeletedAt = field.NewField(tableName, "deleted_at")

	_userCompany.fillFieldMap()

	return _userCompany
}

type userCompany struct {
	userCompanyDo userCompanyDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键id
	UserID       field.String // 用户id
	CompanyID    field.String // 企业id
	InviteUserID field.String // 邀请人id
	InviteTime   field.Time   // 邀请时间
	Status       field.Int32  // 状态：1-正常，2-已禁用
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field

	fieldMap map[string]field.Expr
}

func (u userCompany) Table(newTableName string) *userCompany {
	u.userCompanyDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userCompany) As(alias string) *userCompany {
	u.userCompanyDo.DO = *(u.userCompanyDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userCompany) updateTableName(table string) *userCompany {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UserID = field.NewString(table, "user_id")
	u.CompanyID = field.NewString(table, "company_id")
	u.InviteUserID = field.NewString(table, "invite_user_id")
	u.InviteTime = field.NewTime(table, "invite_time")
	u.Status = field.NewInt32(table, "status")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userCompany) WithContext(ctx context.Context) *userCompanyDo {
	return u.userCompanyDo.WithContext(ctx)
}

func (u userCompany) TableName() string { return u.userCompanyDo.TableName() }

func (u userCompany) Alias() string { return u.userCompanyDo.Alias() }

func (u *userCompany) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userCompany) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 9)
	u.fieldMap["id"] = u.ID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["company_id"] = u.CompanyID
	u.fieldMap["invite_user_id"] = u.InviteUserID
	u.fieldMap["invite_time"] = u.InviteTime
	u.fieldMap["status"] = u.Status
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userCompany) clone(db *gorm.DB) userCompany {
	u.userCompanyDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userCompany) replaceDB(db *gorm.DB) userCompany {
	u.userCompanyDo.ReplaceDB(db)
	return u
}

type userCompanyDo struct{ gen.DO }

func (u userCompanyDo) Debug() *userCompanyDo {
	return u.withDO(u.DO.Debug())
}

func (u userCompanyDo) WithContext(ctx context.Context) *userCompanyDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userCompanyDo) ReadDB() *userCompanyDo {
	return u.Clauses(dbresolver.Read)
}

func (u userCompanyDo) WriteDB() *userCompanyDo {
	return u.Clauses(dbresolver.Write)
}

func (u userCompanyDo) Session(config *gorm.Session) *userCompanyDo {
	return u.withDO(u.DO.Session(config))
}

func (u userCompanyDo) Clauses(conds ...clause.Expression) *userCompanyDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userCompanyDo) Returning(value interface{}, columns ...string) *userCompanyDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userCompanyDo) Not(conds ...gen.Condition) *userCompanyDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userCompanyDo) Or(conds ...gen.Condition) *userCompanyDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userCompanyDo) Select(conds ...field.Expr) *userCompanyDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userCompanyDo) Where(conds ...gen.Condition) *userCompanyDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userCompanyDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userCompanyDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userCompanyDo) Order(conds ...field.Expr) *userCompanyDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userCompanyDo) Distinct(cols ...field.Expr) *userCompanyDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userCompanyDo) Omit(cols ...field.Expr) *userCompanyDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userCompanyDo) Join(table schema.Tabler, on ...field.Expr) *userCompanyDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userCompanyDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userCompanyDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userCompanyDo) RightJoin(table schema.Tabler, on ...field.Expr) *userCompanyDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userCompanyDo) Group(cols ...field.Expr) *userCompanyDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userCompanyDo) Having(conds ...gen.Condition) *userCompanyDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userCompanyDo) Limit(limit int) *userCompanyDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userCompanyDo) Offset(offset int) *userCompanyDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userCompanyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userCompanyDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userCompanyDo) Unscoped() *userCompanyDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userCompanyDo) Create(values ...*model.UserCompany) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userCompanyDo) CreateInBatches(values []*model.UserCompany, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userCompanyDo) Save(values ...*model.UserCompany) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userCompanyDo) First() (*model.UserCompany, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCompany), nil
	}
}

func (u userCompanyDo) Take() (*model.UserCompany, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCompany), nil
	}
}

func (u userCompanyDo) Last() (*model.UserCompany, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCompany), nil
	}
}

func (u userCompanyDo) Find() ([]*model.UserCompany, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserCompany), err
}

func (u userCompanyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserCompany, err error) {
	buf := make([]*model.UserCompany, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userCompanyDo) FindInBatches(result *[]*model.UserCompany, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userCompanyDo) Attrs(attrs ...field.AssignExpr) *userCompanyDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userCompanyDo) Assign(attrs ...field.AssignExpr) *userCompanyDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userCompanyDo) Joins(fields ...field.RelationField) *userCompanyDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userCompanyDo) Preload(fields ...field.RelationField) *userCompanyDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userCompanyDo) FirstOrInit() (*model.UserCompany, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCompany), nil
	}
}

func (u userCompanyDo) FirstOrCreate() (*model.UserCompany, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCompany), nil
	}
}

func (u userCompanyDo) FindByPage(offset int, limit int) (result []*model.UserCompany, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userCompanyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userCompanyDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userCompanyDo) Delete(models ...*model.UserCompany) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userCompanyDo) withDO(do gen.Dao) *userCompanyDo {
	u.DO = *do.(*gen.DO)
	return u
}
