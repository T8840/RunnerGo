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

func newSmsLog(db *gorm.DB, opts ...gen.DOOption) smsLog {
	_smsLog := smsLog{}

	_smsLog.smsLogDo.UseDB(db, opts...)
	_smsLog.smsLogDo.UseModel(&model.SmsLog{})

	tableName := _smsLog.smsLogDo.TableName()
	_smsLog.ALL = field.NewAsterisk(tableName)
	_smsLog.ID = field.NewInt64(tableName, "id")
	_smsLog.Type = field.NewInt32(tableName, "type")
	_smsLog.Mobile = field.NewString(tableName, "mobile")
	_smsLog.Content = field.NewString(tableName, "content")
	_smsLog.VerifyCode = field.NewString(tableName, "verify_code")
	_smsLog.VerifyCodeExpirationTime = field.NewTime(tableName, "verify_code_expiration_time")
	_smsLog.ClientIP = field.NewString(tableName, "client_ip")
	_smsLog.SendStatus = field.NewInt32(tableName, "send_status")
	_smsLog.VerifyStatus = field.NewInt32(tableName, "verify_status")
	_smsLog.SendResponse = field.NewString(tableName, "send_response")
	_smsLog.CreatedAt = field.NewTime(tableName, "created_at")
	_smsLog.UpdatedAt = field.NewTime(tableName, "updated_at")
	_smsLog.DeletedAt = field.NewField(tableName, "deleted_at")

	_smsLog.fillFieldMap()

	return _smsLog
}

type smsLog struct {
	smsLogDo smsLogDo

	ALL                      field.Asterisk
	ID                       field.Int64  // 主键id
	Type                     field.Int32  // 短信类型: 1-注册，2-登录，3-找回密码
	Mobile                   field.String // 手机号
	Content                  field.String // 短信内容
	VerifyCode               field.String // 验证码
	VerifyCodeExpirationTime field.Time   // 验证码有效时间
	ClientIP                 field.String // 客户端IP
	SendStatus               field.Int32  // 发送状态：1-成功 2-失败
	VerifyStatus             field.Int32  // 校验状态：1-未校验 2-已校验
	SendResponse             field.String // 短信服务响应
	CreatedAt                field.Time   // 创建时间
	UpdatedAt                field.Time   // 修改时间
	DeletedAt                field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (s smsLog) Table(newTableName string) *smsLog {
	s.smsLogDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsLog) As(alias string) *smsLog {
	s.smsLogDo.DO = *(s.smsLogDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsLog) updateTableName(table string) *smsLog {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Type = field.NewInt32(table, "type")
	s.Mobile = field.NewString(table, "mobile")
	s.Content = field.NewString(table, "content")
	s.VerifyCode = field.NewString(table, "verify_code")
	s.VerifyCodeExpirationTime = field.NewTime(table, "verify_code_expiration_time")
	s.ClientIP = field.NewString(table, "client_ip")
	s.SendStatus = field.NewInt32(table, "send_status")
	s.VerifyStatus = field.NewInt32(table, "verify_status")
	s.SendResponse = field.NewString(table, "send_response")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *smsLog) WithContext(ctx context.Context) *smsLogDo { return s.smsLogDo.WithContext(ctx) }

func (s smsLog) TableName() string { return s.smsLogDo.TableName() }

func (s smsLog) Alias() string { return s.smsLogDo.Alias() }

func (s *smsLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsLog) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 13)
	s.fieldMap["id"] = s.ID
	s.fieldMap["type"] = s.Type
	s.fieldMap["mobile"] = s.Mobile
	s.fieldMap["content"] = s.Content
	s.fieldMap["verify_code"] = s.VerifyCode
	s.fieldMap["verify_code_expiration_time"] = s.VerifyCodeExpirationTime
	s.fieldMap["client_ip"] = s.ClientIP
	s.fieldMap["send_status"] = s.SendStatus
	s.fieldMap["verify_status"] = s.VerifyStatus
	s.fieldMap["send_response"] = s.SendResponse
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s smsLog) clone(db *gorm.DB) smsLog {
	s.smsLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s smsLog) replaceDB(db *gorm.DB) smsLog {
	s.smsLogDo.ReplaceDB(db)
	return s
}

type smsLogDo struct{ gen.DO }

func (s smsLogDo) Debug() *smsLogDo {
	return s.withDO(s.DO.Debug())
}

func (s smsLogDo) WithContext(ctx context.Context) *smsLogDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsLogDo) ReadDB() *smsLogDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsLogDo) WriteDB() *smsLogDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsLogDo) Session(config *gorm.Session) *smsLogDo {
	return s.withDO(s.DO.Session(config))
}

func (s smsLogDo) Clauses(conds ...clause.Expression) *smsLogDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsLogDo) Returning(value interface{}, columns ...string) *smsLogDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsLogDo) Not(conds ...gen.Condition) *smsLogDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsLogDo) Or(conds ...gen.Condition) *smsLogDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsLogDo) Select(conds ...field.Expr) *smsLogDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsLogDo) Where(conds ...gen.Condition) *smsLogDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsLogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *smsLogDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s smsLogDo) Order(conds ...field.Expr) *smsLogDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsLogDo) Distinct(cols ...field.Expr) *smsLogDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsLogDo) Omit(cols ...field.Expr) *smsLogDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsLogDo) Join(table schema.Tabler, on ...field.Expr) *smsLogDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *smsLogDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *smsLogDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsLogDo) Group(cols ...field.Expr) *smsLogDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsLogDo) Having(conds ...gen.Condition) *smsLogDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsLogDo) Limit(limit int) *smsLogDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsLogDo) Offset(offset int) *smsLogDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *smsLogDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsLogDo) Unscoped() *smsLogDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsLogDo) Create(values ...*model.SmsLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsLogDo) CreateInBatches(values []*model.SmsLog, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsLogDo) Save(values ...*model.SmsLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsLogDo) First() (*model.SmsLog, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsLog), nil
	}
}

func (s smsLogDo) Take() (*model.SmsLog, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsLog), nil
	}
}

func (s smsLogDo) Last() (*model.SmsLog, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsLog), nil
	}
}

func (s smsLogDo) Find() ([]*model.SmsLog, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsLog), err
}

func (s smsLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsLog, err error) {
	buf := make([]*model.SmsLog, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsLogDo) FindInBatches(result *[]*model.SmsLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsLogDo) Attrs(attrs ...field.AssignExpr) *smsLogDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsLogDo) Assign(attrs ...field.AssignExpr) *smsLogDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsLogDo) Joins(fields ...field.RelationField) *smsLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsLogDo) Preload(fields ...field.RelationField) *smsLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsLogDo) FirstOrInit() (*model.SmsLog, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsLog), nil
	}
}

func (s smsLogDo) FirstOrCreate() (*model.SmsLog, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsLog), nil
	}
}

func (s smsLogDo) FindByPage(offset int, limit int) (result []*model.SmsLog, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s smsLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsLogDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsLogDo) Delete(models ...*model.SmsLog) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsLogDo) withDO(do gen.Dao) *smsLogDo {
	s.DO = *do.(*gen.DO)
	return s
}
