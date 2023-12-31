package consts

const (
	TeamTypePrivate = 1 // 私有团队
	TeamTypeNormal  = 2 // 普通团队

	PrivateTeamTrialTime  = 360  //私有团队试用期7天 单位:h
	PrivateTeamDefaultVum = 5000 //私有团队试用期默认赠送vum数

	IsVipTypeNo  = 1 // 否
	IsVipTypeYes = 2 // 是

	// 个人版试用期内赠送成员数
	TrialTeamGiftUserNum = 1

	// 所有团队类型的最大并发数，固定是50000
	AllTeamMaxConcurrency = 50000

	// 团队是否可用
	TeamNotCanUse = 0 // 不可用
	TeamCanUse    = 1 // 可用

	TeamListOrderUserID     = 1 //  1:我加入的团队  2:我是团队管理员的团队  3:我收藏的团队
	TeamListOrderSuper      = 2 //
	TeamListOrderCollection = 3 //
	TeamListOrderCompany    = 4 // 企业下的所有团队

	TeamCollection   = 1 // 收藏
	TeamUnCollection = 2 // 取消收藏

	TeamIsShow      = 1 // 是否展示在列表中 1：展示   2：不展示
	TeamIsShowFalse = 2 //
)
