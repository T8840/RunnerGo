// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                     db,
		Company:                newCompany(db, opts...),
		Permission:             newPermission(db, opts...),
		PublicFunction:         newPublicFunction(db, opts...),
		Role:                   newRole(db, opts...),
		RolePermission:         newRolePermission(db, opts...),
		RoleTypePermission:     newRoleTypePermission(db, opts...),
		Setting:                newSetting(db, opts...),
		Team:                   newTeam(db, opts...),
		ThirdNotice:            newThirdNotice(db, opts...),
		ThirdNoticeChannel:     newThirdNoticeChannel(db, opts...),
		ThirdNoticeGroup:       newThirdNoticeGroup(db, opts...),
		ThirdNoticeGroupRelate: newThirdNoticeGroupRelate(db, opts...),
		User:                   newUser(db, opts...),
		UserCompany:            newUserCompany(db, opts...),
		UserRole:               newUserRole(db, opts...),
		UserTeam:               newUserTeam(db, opts...),
		UserTeamCollection:     newUserTeamCollection(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Company                company
	Permission             permission
	PublicFunction         publicFunction
	Role                   role
	RolePermission         rolePermission
	RoleTypePermission     roleTypePermission
	Setting                setting
	Team                   team
	ThirdNotice            thirdNotice
	ThirdNoticeChannel     thirdNoticeChannel
	ThirdNoticeGroup       thirdNoticeGroup
	ThirdNoticeGroupRelate thirdNoticeGroupRelate
	User                   user
	UserCompany            userCompany
	UserRole               userRole
	UserTeam               userTeam
	UserTeamCollection     userTeamCollection
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                     db,
		Company:                q.Company.clone(db),
		Permission:             q.Permission.clone(db),
		PublicFunction:         q.PublicFunction.clone(db),
		Role:                   q.Role.clone(db),
		RolePermission:         q.RolePermission.clone(db),
		RoleTypePermission:     q.RoleTypePermission.clone(db),
		Setting:                q.Setting.clone(db),
		Team:                   q.Team.clone(db),
		ThirdNotice:            q.ThirdNotice.clone(db),
		ThirdNoticeChannel:     q.ThirdNoticeChannel.clone(db),
		ThirdNoticeGroup:       q.ThirdNoticeGroup.clone(db),
		ThirdNoticeGroupRelate: q.ThirdNoticeGroupRelate.clone(db),
		User:                   q.User.clone(db),
		UserCompany:            q.UserCompany.clone(db),
		UserRole:               q.UserRole.clone(db),
		UserTeam:               q.UserTeam.clone(db),
		UserTeamCollection:     q.UserTeamCollection.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                     db,
		Company:                q.Company.replaceDB(db),
		Permission:             q.Permission.replaceDB(db),
		PublicFunction:         q.PublicFunction.replaceDB(db),
		Role:                   q.Role.replaceDB(db),
		RolePermission:         q.RolePermission.replaceDB(db),
		RoleTypePermission:     q.RoleTypePermission.replaceDB(db),
		Setting:                q.Setting.replaceDB(db),
		Team:                   q.Team.replaceDB(db),
		ThirdNotice:            q.ThirdNotice.replaceDB(db),
		ThirdNoticeChannel:     q.ThirdNoticeChannel.replaceDB(db),
		ThirdNoticeGroup:       q.ThirdNoticeGroup.replaceDB(db),
		ThirdNoticeGroupRelate: q.ThirdNoticeGroupRelate.replaceDB(db),
		User:                   q.User.replaceDB(db),
		UserCompany:            q.UserCompany.replaceDB(db),
		UserRole:               q.UserRole.replaceDB(db),
		UserTeam:               q.UserTeam.replaceDB(db),
		UserTeamCollection:     q.UserTeamCollection.replaceDB(db),
	}
}

type queryCtx struct {
	Company                *companyDo
	Permission             *permissionDo
	PublicFunction         *publicFunctionDo
	Role                   *roleDo
	RolePermission         *rolePermissionDo
	RoleTypePermission     *roleTypePermissionDo
	Setting                *settingDo
	Team                   *teamDo
	ThirdNotice            *thirdNoticeDo
	ThirdNoticeChannel     *thirdNoticeChannelDo
	ThirdNoticeGroup       *thirdNoticeGroupDo
	ThirdNoticeGroupRelate *thirdNoticeGroupRelateDo
	User                   *userDo
	UserCompany            *userCompanyDo
	UserRole               *userRoleDo
	UserTeam               *userTeamDo
	UserTeamCollection     *userTeamCollectionDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Company:                q.Company.WithContext(ctx),
		Permission:             q.Permission.WithContext(ctx),
		PublicFunction:         q.PublicFunction.WithContext(ctx),
		Role:                   q.Role.WithContext(ctx),
		RolePermission:         q.RolePermission.WithContext(ctx),
		RoleTypePermission:     q.RoleTypePermission.WithContext(ctx),
		Setting:                q.Setting.WithContext(ctx),
		Team:                   q.Team.WithContext(ctx),
		ThirdNotice:            q.ThirdNotice.WithContext(ctx),
		ThirdNoticeChannel:     q.ThirdNoticeChannel.WithContext(ctx),
		ThirdNoticeGroup:       q.ThirdNoticeGroup.WithContext(ctx),
		ThirdNoticeGroupRelate: q.ThirdNoticeGroupRelate.WithContext(ctx),
		User:                   q.User.WithContext(ctx),
		UserCompany:            q.UserCompany.WithContext(ctx),
		UserRole:               q.UserRole.WithContext(ctx),
		UserTeam:               q.UserTeam.WithContext(ctx),
		UserTeamCollection:     q.UserTeamCollection.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
