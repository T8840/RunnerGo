package team

import (
	"context"
	"fmt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/conf"
	"github.com/gin-gonic/gin"
	"github.com/go-omnibus/proof"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"

	"github.com/go-omnibus/omnibus"

	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/encrypt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/mail"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/model"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/query"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/packer"
)

func SaveTeam(ctx *gin.Context, req *rao.SaveTeamReq) (string, error) {
	err := dal.GetQuery().Transaction(func(tx *query.Query) error {
		_, err := tx.Team.WithContext(ctx).Where(tx.Team.TeamID.Eq(req.TeamID)).UpdateColumnSimple(tx.Team.Name.Value(req.Name))
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return "", err
	}

	return req.TeamID, nil
}

func ListByUserID(ctx context.Context, userID string) ([]*rao.Team, error) {

	ut := query.Use(dal.DB()).UserTeam
	userTeams, err := ut.WithContext(ctx).Where(ut.UserID.Eq(userID)).Order(ut.CreatedAt.Desc()).Find()
	if err != nil {
		return nil, err
	}

	var teamIDs []string
	for _, team := range userTeams {
		teamIDs = append(teamIDs, team.TeamID)
	}

	t := query.Use(dal.DB()).Team
	teams, err := t.WithContext(ctx).Where(t.TeamID.In(teamIDs...)).Find()
	if err != nil {
		return nil, err
	}

	var NewTeamIDs []string
	for _, team := range teams {
		NewTeamIDs = append(NewTeamIDs, team.TeamID)
	}

	// 再扯查询可用的用户团队数据
	userTeamsNew, err := ut.WithContext(ctx).Where(ut.UserID.Eq(userID), ut.TeamID.In(NewTeamIDs...)).Order(ut.CreatedAt.Desc()).Find()
	if err != nil {
		return nil, err
	}

	var teamCnt []*packer.TeamMemberCount
	if err := ut.WithContext(ctx).Select(ut.TeamID, ut.UserID.Count().As("cnt")).Where(ut.TeamID.In(teamIDs...)).Group(ut.TeamID).Scan(&teamCnt); err != nil {
		return nil, err
	}

	var userIDs []string
	for _, team := range teams {
		userIDs = append(userIDs, team.CreatedUserID)
	}
	u := dal.GetQuery().User
	users, err := u.WithContext(ctx).Where(u.UserID.In(userIDs...)).Find()
	if err != nil {
		return nil, err
	}

	return packer.TransTeamsModelToRaoTeam(teams, userTeamsNew, teamCnt, users), nil
}

func ListMembersByTeamID(ctx context.Context, teamID string) ([]*rao.Member, error) {
	ut := query.Use(dal.DB()).UserTeam
	userTeams, err := ut.WithContext(ctx).Where(ut.TeamID.Eq(teamID)).Order(ut.RoleID).Find()
	if err != nil {
		return nil, err
	}

	var userIDs []string
	for _, team := range userTeams {
		userIDs = append(userIDs, team.UserID)
		userIDs = append(userIDs, team.InviteUserID)
	}

	u := query.Use(dal.DB()).User
	users, err := u.WithContext(ctx).Where(u.UserID.In(userIDs...)).Find()
	if err != nil {
		return nil, err
	}

	return packer.TransUsersToRaoMembers(users, userTeams), nil
}

func InviteMember(ctx context.Context, inviteUserID string, teamID string, members []*rao.InviteMember) (*rao.InviteMemberResp, error) {

	var emails []string
	memo := make(map[string]int64)
	for _, member := range members {
		emails = append(emails, member.Email)
		memo[member.Email] = member.RoleID
	}

	tx := query.Use(dal.DB()).User
	users, err := tx.WithContext(ctx).Where(tx.Email.In(emails...)).Find()
	if err != nil {
		return nil, err
	}

	var registerEmail []string
	for _, user := range users {
		registerEmail = append(registerEmail, user.Email)
	}
	registerEmail = omnibus.StringArrayUnique(registerEmail)

	var userIDs []string
	for _, user := range users {
		userIDs = append(userIDs, user.UserID)
	}
	utt := dal.GetQuery().UserTeam
	existUser, err := utt.WithContext(ctx).Where(utt.TeamID.Eq(teamID), utt.UserID.In(userIDs...)).Find()
	if err != nil {
		return nil, err
	}

	for i, user := range users {
		for _, eu := range existUser {
			if eu.UserID == user.UserID {
				users[i] = nil
			}
		}
	}

	var ut []*model.UserTeam
	for _, user := range users {
		if user != nil {
			ut = append(ut, &model.UserTeam{
				UserID:       user.UserID,
				TeamID:       teamID,
				InviteUserID: inviteUserID,
				RoleID:       memo[user.Email],
			})
		}
	}

	if err := query.Use(dal.DB()).UserTeam.WithContext(ctx).CreateInBatches(ut, 5); err != nil {
		return nil, err
	}

	u, err := tx.WithContext(ctx).Where(tx.UserID.Eq(inviteUserID)).First()
	if err != nil {
		return nil, err
	}

	px := dal.GetQuery().Team
	t, err := px.WithContext(ctx).Where(px.TeamID.Eq(teamID)).First()
	if err != nil {
		return nil, err
	}

	for _, e := range registerEmail {
		if err := mail.SendInviteEmail(e, inviteUserID, u.Nickname, t.Name, teamID, memo[e], true); err != nil {
			return nil, err
		}
	}

	unRegisterEmail := omnibus.StringArrayUnique(omnibus.StringArrayDiff(emails, registerEmail))
	if len(unRegisterEmail) > 0 {
		var userQueue []*model.TeamUserQueue
		for _, e := range unRegisterEmail {
			if err := mail.SendInviteEmail(e, inviteUserID, u.Nickname, t.Name, teamID, memo[e], false); err != nil {
				return nil, err
			}

			userQueue = append(userQueue, &model.TeamUserQueue{
				Email:  e,
				TeamID: teamID,
			})
		}
		qx := dal.GetQuery().TeamUserQueue
		if err := qx.WithContext(ctx).CreateInBatches(userQueue, 5); err != nil {
			return nil, err
		}
	}

	return &rao.InviteMemberResp{
		RegisterNum:      len(registerEmail),
		UnRegisterNum:    len(unRegisterEmail),
		UnRegisterEmails: unRegisterEmail,
	}, nil
}

func SetTeamRole(ctx context.Context, teamID string, userID string, roleID int64) error {
	return dal.GetQuery().Transaction(func(tx *query.Query) error {
		_, err := tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Eq(teamID), tx.UserTeam.UserID.Eq(userID)).First()
		if err != nil {
			return err
		}

		_, err = tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Eq(teamID), tx.UserTeam.UserID.Eq(userID)).UpdateColumn(tx.UserTeam.RoleID, roleID)

		return err
	})
}

func RemoveMember(ctx *gin.Context, teamID string, userID string, memberID string) error {
	newTeamID := ""
	return dal.GetQuery().Transaction(func(tx *query.Query) error {
		// 不能移除自己
		if userID == memberID {
			return fmt.Errorf("user no permissions")
		}

		admin, err := tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Eq(teamID), tx.UserTeam.UserID.Eq(userID)).First()
		if err != nil {
			return err
		}

		// 只有管理员和创建人可以操作移除
		if !omnibus.InArray(admin.RoleID, []int64{consts.RoleTypeAdmin, consts.RoleTypeOwner}) {
			return fmt.Errorf("user no permissions")
		}

		user, err := tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Eq(teamID), tx.UserTeam.UserID.Eq(memberID)).First()
		if err != nil {
			return err
		}

		// 不能移除创建人
		if user.RoleID == consts.RoleTypeOwner {
			return fmt.Errorf("user no permissions")
		}

		// 只有创建人能移除管理员
		if user.RoleID == consts.RoleTypeAdmin && admin.RoleID != consts.RoleTypeOwner {
			return fmt.Errorf("user no permissions")
		}

		_, err = tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Eq(teamID), tx.UserTeam.UserID.Eq(memberID)).Delete()

		// 查询被移除用户的默认团队是否是被移除的团队
		memberDefaultTeamInfo, err := tx.Setting.WithContext(ctx).Where(tx.Setting.UserID.Eq(memberID)).First()
		if err != nil {
			return fmt.Errorf("没有找到当前用户的默认团队")
		}

		// 如果被移除成员的默认团队是被移除出去的团队
		if memberDefaultTeamInfo.TeamID == teamID {
			// 修改移除成员到新地默认团队
			newTeamID = GetNewDefaultTeamID(ctx, teamID, userID)
			// 修改用户的默认团队为新的团队id
			_, err = tx.Setting.WithContext(ctx).Where(tx.Setting.UserID.Eq(memberID)).UpdateColumnSimple(tx.Setting.TeamID.Value(newTeamID))
			if err != nil {
				return err
			}
		}
		return err
	})
}

func QuitTeam(ctx *gin.Context, teamID string, userID string) (string, error) {
	newTeamID := ""
	err := dal.GetQuery().Transaction(func(tx *query.Query) error {
		team, err := tx.Team.WithContext(ctx).Where(tx.Team.TeamID.Eq(teamID)).First()
		if err != nil {
			return err
		}

		// 不能退出私有团队
		if team.CreatedUserID == userID {
			return fmt.Errorf("user no permissions")
		}

		ut, err := tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Eq(teamID), tx.UserTeam.UserID.Eq(userID)).First()
		if err != nil {
			return err
		}

		switch ut.RoleID {
		case consts.RoleTypeOwner:
			cnt, err := tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Eq(teamID), tx.UserTeam.RoleID.Eq(consts.RoleTypeAdmin)).Count()
			if err != nil {
				return err
			}
			if cnt == 0 {
				return fmt.Errorf("not found admin user")
			}
		case consts.RoleTypeMember, consts.RoleTypeAdmin:
			break
		}

		_, err = tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Eq(teamID), tx.UserTeam.UserID.Eq(userID)).Delete()
		if err != nil {
			return err
		}

		newTeamID = GetNewDefaultTeamID(ctx, teamID, userID)

		// 修改用户的默认团队为新的团队id
		//_, err = tx.Setting.WithContext(ctx).Where(tx.Setting.UserID.Eq(userID)).UpdateColumnSimple(tx.Setting.TeamID.Value(newTeamID))
		//if err != nil {
		//	return err
		//}

		return nil
	})

	if err != nil {
		return "", err
	}
	return newTeamID, nil
}

func DisbandTeam(ctx context.Context, teamID string, userID string) (string, error) {
	newTeamID := ""
	err := dal.GetQuery().Transaction(func(tx *query.Query) error {
		t, err := tx.Team.WithContext(ctx).Where(
			tx.Team.TeamID.Eq(teamID), tx.Team.Type.Eq(consts.TeamTypeNormal), tx.Team.CreatedUserID.Eq(userID)).First()
		if err != nil {
			return err
		}

		//之前逻辑--解散团队以后，把默认团队设置为各自的付费团队--如果没有，则设置为0
		settings, err := tx.Setting.WithContext(ctx).Where(tx.Setting.TeamID.Eq(teamID)).Find()
		if err != nil {
			return err
		}

		for _, s := range settings {
			var otherUserNewTeamID string = ""
			// 查询当前用户是否有付费团队
			utList, err := tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Neq(teamID), tx.UserTeam.UserID.Eq(s.UserID)).Find()
			userAllTeamIDs := make([]string, 0, len(utList))
			if err == nil && len(utList) > 0 {
				for _, utInfo := range utList {
					userAllTeamIDs = append(userAllTeamIDs, utInfo.TeamID)
				}
			}

			nowTime := time.Now().Unix()
			teamList, err := tx.Team.WithContext(ctx).Where(tx.Team.TeamID.In(userAllTeamIDs...)).Find()
			if err == nil && len(teamList) > 0 {
				for _, teamInfo := range teamList {
					if teamInfo.VipExpirationDate.Unix() > nowTime {
						if s.UserID == userID {
							newTeamID = teamInfo.TeamID
						} else {
							otherUserNewTeamID = teamInfo.TeamID
						}
						break
					}
				}
			}

			if s.UserID != userID {
				_, err = tx.Setting.WithContext(ctx).Where(tx.Setting.ID.Eq(s.ID)).UpdateColumn(tx.Setting.TeamID, otherUserNewTeamID)
				if err != nil {
					return err
				}
			}
		}

		// 删除当前团队
		_, err = tx.Team.WithContext(ctx).Where(tx.Team.TeamID.Eq(t.TeamID)).Delete()
		if err != nil {
			return err
		}

		// 删除所有用户与解散团队之间的关系数据
		_, err = tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Eq(t.TeamID)).Delete()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return newTeamID, err
}

func GetNewDefaultTeamID(ctx *gin.Context, teamID string, userID string) string {
	newTeamID := ""
	// 查询当前用户是否有付费团队
	tx := dal.GetQuery()
	utList, err := tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Neq(teamID), tx.UserTeam.UserID.Eq(userID)).Find()
	//userAllTeamIDs := make([]string, 0, len(utList))
	if err == nil && len(utList) > 0 {
		for _, utInfo := range utList {
			//userAllTeamIDs = append(userAllTeamIDs, utInfo.TeamID)
			newTeamID = utInfo.TeamID
		}
	}

	//nowTime := time.Now().Unix()
	//teamList, err := tx.Team.WithContext(ctx).Where(tx.Team.TeamID.In(userAllTeamIDs...)).Find()
	//if err == nil && len(teamList) > 0 {
	//	for _, teamInfo := range teamList {
	//		if teamInfo.VipExpirationDate.Unix() > nowTime {
	//			newTeamID = teamInfo.TeamID
	//			break
	//		}
	//	}
	//}
	return newTeamID
}

func GetInviteUserInfo(ctx *gin.Context, req *rao.GetInviteUserInfoReq) (*rao.GetInviteUserInfoResp, error) {
	userInfoString := encrypt.AesDecrypt(req.InviteVerifyCode, conf.Conf.InviteData.AesSecretKey)
	userInfoArr := strings.Split(userInfoString, "_")
	if len(userInfoArr) != 4 {
		return nil, fmt.Errorf("验证码解析错误")
	}

	teamID := userInfoArr[0]
	roleID, _ := strconv.ParseInt(userInfoArr[1], 10, 64)
	inviteUserID := userInfoArr[2]

	// 检查邀请链接是否过期
	k := fmt.Sprintf("invite:url:%s:%d:%s:%s", teamID, roleID, inviteUserID, req.InviteVerifyCode)
	_, err := dal.GetRDB().Get(ctx, k).Result()
	if err != nil {
		return nil, fmt.Errorf("邀请链接已过期")
	}

	tx := dal.GetQuery()
	// 查询团队信息
	teamInfo, err := tx.Team.WithContext(ctx).Where(tx.Team.TeamID.Eq(teamID)).First()
	if err != nil {
		return nil, fmt.Errorf("团队不存在或已解散")
	}

	// 获取邀请人信息
	inviteUserInfo, err := tx.User.WithContext(ctx).Where(tx.User.UserID.Eq(inviteUserID)).First()
	if err != nil {
		return nil, err
	}

	// 查询当前团队已存在用户信息
	existTeamUserCount, err := tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Eq(teamID)).Count()
	if err != nil {
		return nil, err
	}

	if existTeamUserCount >= teamInfo.MaxUserNum {
		return nil, fmt.Errorf("团队人数已超上限")
	}

	res := &rao.GetInviteUserInfoResp{
		TeamID:         teamID,
		RoleID:         roleID,
		InviteUserID:   inviteUserID,
		InviteUserName: inviteUserInfo.Nickname,
		TeamName:       teamInfo.Name,
	}
	return res, nil
}

func InviteLogin(ctx *gin.Context, verifyCode string, userID string) error {
	userInfoString := encrypt.AesDecrypt(verifyCode, conf.Conf.InviteData.AesSecretKey)
	userInfoArr := strings.Split(userInfoString, "_")
	if len(userInfoArr) != 4 {
		return fmt.Errorf("验证码解析错误")
	}

	teamID := userInfoArr[0]
	roleID, _ := strconv.ParseInt(userInfoArr[1], 10, 64)
	inviteUserID := userInfoArr[2]

	// 把当前用户的当前团队设置为被邀请团队
	err := query.Use(dal.DB()).Transaction(func(tx *query.Query) error {
		// 1、把用户当前所属团队修改为被邀请的团队
		updateData := make(map[string]interface{}, 1)
		updateData["team_id"] = teamID
		_, err := tx.Setting.WithContext(ctx).Where(tx.Setting.UserID.Eq(userID)).Updates(updateData)
		if err != nil {
			proof.Infof("邀请登录--修改用户当前团队失败，err:", err)
			return err
		}
		// 2、把当前用户放到被邀请的团队里面
		_, err = tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.UserID.Eq(userID)).Where(tx.UserTeam.TeamID.Eq(teamID)).First()
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if err == gorm.ErrRecordNotFound { // 没查到，就插入
			insertData := &model.UserTeam{
				UserID:       userID,
				TeamID:       teamID,
				RoleID:       roleID,
				InviteUserID: inviteUserID,
			}
			err = tx.UserTeam.WithContext(ctx).Create(insertData)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetInviteEmailIsExist(ctx *gin.Context, req *rao.GetInviteEmailIsExistReq) (bool, error) {
	// 查询团队信息
	tx := dal.GetQuery()
	teamInfo, err := tx.Team.WithContext(ctx).Where(tx.Team.TeamID.Eq(req.TeamID)).First()
	if err != nil {
		return false, err
	}

	// 当前团队已存在多少人
	allTeamUserNum, err := tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Eq(req.TeamID)).Count()
	if err != nil {
		return false, err
	}

	if allTeamUserNum >= teamInfo.MaxUserNum {
		return false, fmt.Errorf("受邀人数已超出团队成员限额")
	}

	// 查询当前邮箱所属用户是否已存在
	userInfo, err := tx.User.WithContext(ctx).Where(tx.User.Email.Eq(req.Email)).First()
	if err == nil { // 找到了
		_, err = tx.UserTeam.WithContext(ctx).Where(tx.UserTeam.TeamID.Eq(req.TeamID),
			tx.UserTeam.UserID.Eq(userInfo.UserID)).First()
		if err == nil {
			return true, nil
		}
	}
	return false, nil
}

// GetTeamIsOverdue 判断当前团队是否过期
func GetTeamIsOverdue(ctx context.Context, teamID string) (bool, error) {
	res := false
	// 查询当前团队是否过期
	tx := dal.GetQuery().Team
	teamInfo, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(teamID)).First()
	if err == nil {
		nowTime := time.Now().Unix()
		teamExpirationDate := teamInfo.VipExpirationDate.Unix()
		if teamExpirationDate < nowTime {
			res = true
		}
	}
	return res, err
}
