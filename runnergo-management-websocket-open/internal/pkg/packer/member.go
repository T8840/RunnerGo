package packer

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/model"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
)

func TransUsersToRaoMembers(users []*model.User, userTeams []*model.UserTeam) []*rao.Member {
	ret := make([]*rao.Member, 0)

	memo := make(map[string]string)
	for _, u := range users {
		memo[u.UserID] = u.Nickname
	}

	for _, ut := range userTeams {
		for _, u := range users {
			if ut.UserID == u.UserID {
				ret = append(ret, &rao.Member{
					Avatar:         u.Avatar,
					UserID:         u.UserID,
					Email:          u.Email,
					Nickname:       u.Nickname,
					JoinTimeSec:    ut.CreatedAt.Unix(),
					RoleID:         ut.RoleID,
					InviteUserID:   ut.InviteUserID,
					InviteUserName: memo[ut.InviteUserID],
				})
			}
		}
	}

	//for _, u := range users {
	//	for _, ut := range userTeams {
	//		if ut.UserID == u.ID {
	//
	//			ret = append(ret, &rao.Member{
	//				Avatar:         u.Avatar,
	//				UserID:         u.ID,
	//				Email:          u.Email,
	//				Nickname:       u.Nickname,
	//				JoinTimeSec:    ut.CreatedAt.Unix(),
	//				RoleID:         ut.RoleID,
	//				InviteUserID:   ut.InviteUserID,
	//				InviteUserName: memo[ut.InviteUserID],
	//			})
	//		}
	//	}
	//}
	return ret
}
