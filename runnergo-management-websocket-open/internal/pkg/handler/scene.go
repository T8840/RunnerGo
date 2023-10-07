package handler

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/errno"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/jwt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/response"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/runner"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/scene"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/target"
	"github.com/gin-gonic/gin"
)

// SendScene 调试场景
func SendScene(ctx *gin.Context) {
	var req rao.SendSceneReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	retID, err := target.SendScene(ctx, req.TeamID, req.SceneID, jwt.GetUserIDByCtx(ctx))
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.SendSceneResp{RetID: retID})
	return
}

// StopScene 停止调试场景
func StopScene(ctx *gin.Context) {
	var req rao.StopSceneReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	err := runner.StopScene(ctx, &req)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrHttpFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}

// SendSceneAPI 场景调试接口
func SendSceneAPI(ctx *gin.Context) {
	var req rao.SendSceneAPIReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	retID, err := target.SendSceneAPI(ctx, req.TeamID, req.SceneID, req.NodeID, req.SceneCaseID)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.SendSceneAPIResp{RetID: retID})
	return
}

// GetSendSceneResult 获取调试场景结果
func GetSendSceneResult(ctx *gin.Context) {
	var req rao.GetSendSceneResultReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	s, err := target.GetSendSceneResult(ctx, req.RetID)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.GetSendSceneResultResp{Scenes: s})
	return
}

// SaveScene 创建/修改场景
func SaveScene(ctx *gin.Context) {
	var req rao.SaveSceneReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	targetID, targetName, err := scene.Save(ctx, &req, jwt.GetUserIDByCtx(ctx))
	if err != nil {
		if err.Error() == "名称已存在" {
			response.ErrorWithMsg(ctx, errno.ErrSceneNameAlreadyExist, err.Error())
		} else {
			response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		}

		return
	}

	response.SuccessWithData(ctx, rao.SaveSceneResp{
		TargetID:   targetID,
		TargetName: targetName,
	})
	return
}

// BatchGetScene 获取场景
func BatchGetScene(ctx *gin.Context) {
	var req rao.GetSceneReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	s, err := scene.BatchGetByTargetID(ctx, req.TeamID, req.TargetID, req.Source)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.GetSceneResp{Scenes: s})
	return
}

// SaveFlow 保存场景流程
func SaveFlow(ctx *gin.Context) {
	var req rao.SaveFlowReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	for _, nodeInfo := range req.Nodes {
		if nodeInfo.Type == consts.FlowTypeWaitController && nodeInfo.WaitMs > 20000 {
			response.ErrorWithMsg(ctx, errno.ErrWaitControllerOvertopLimit, "")
			return
		}
	}

	errNum, err := scene.SaveFlow(ctx, &req)
	if err != nil {
		response.ErrorWithMsg(ctx, errNum, err.Error())
		return
	}
	response.Success(ctx)
	return
}

// GetFlow 获取场景流程
func GetFlow(ctx *gin.Context) {
	var req rao.GetFlowReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	resp, err := scene.GetFlow(ctx, req.SceneID)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, resp)
	return
}

// BatchGetFlow 批量获取场景流程
func BatchGetFlow(ctx *gin.Context) {
	var req rao.BatchGetFlowReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	flows, err := scene.BatchGetFlow(ctx, req.SceneID)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.BatchGetFlowResp{Flows: flows})
	return
}

// DeleteScene 删除计划下的场景
func DeleteScene(ctx *gin.Context) {
	var req rao.DeleteSceneReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	if err := scene.DeleteScene(ctx, &req, jwt.GetUserIDByCtx(ctx)); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}
