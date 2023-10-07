package api

import (
	"context"
	"fmt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/jwt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/log"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/uuid"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/record"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/mao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/query"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/packer"
)

func Save(ctx context.Context, req *rao.SaveTargetReq, userID string) (string, error) {
	target := packer.TransSaveTargetReqToTargetModel(req, userID)
	api := packer.TransSaveTargetReqToMaoAPI(req)

	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectAPI)

	err := query.Use(dal.DB()).Transaction(func(tx *query.Query) error {
		// 接口名排重
		_, err := tx.Target.WithContext(ctx).Where(tx.Target.TeamID.Eq(req.TeamID), tx.Target.Name.Eq(req.Name),
			tx.Target.TargetType.Eq(consts.TargetTypeAPI), tx.Target.TargetID.Neq(req.TargetID),
			tx.Target.Status.Eq(consts.TargetStatusNormal)).First()
		if err == nil {
			return fmt.Errorf("名称已存在")
		}

		// 查询当前接口是否存在
		_, err = tx.Target.WithContext(ctx).Where(tx.Target.TargetID.Eq(req.TargetID)).First()
		if err != nil { // 需新增
			if err := tx.Target.WithContext(ctx).Create(target); err != nil {
				return err
			}
			api.TargetID = target.TargetID
			_, err := collection.InsertOne(ctx, api)
			if err != nil {
				return err
			}
			return record.InsertCreate(ctx, target.TeamID, userID, record.OperationOperateCreateAPI, target.Name)
		}

		if _, err := tx.Target.WithContext(ctx).Where(tx.Target.TargetID.Eq(req.TargetID)).Updates(target); err != nil {
			return err
		}

		_, err = collection.UpdateOne(ctx, bson.D{{"target_id", target.TargetID}}, bson.M{"$set": api})
		if err != nil {
			return err
		}
		return record.InsertUpdate(ctx, target.TeamID, userID, record.OperationOperateUpdateAPI, target.Name)
	})
	return target.TargetID, err
}

func SaveImportApi(ctx *gin.Context, req *rao.SaveImportApiReq, userID string) error {
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectAPI)
	err := query.Use(dal.DB()).Transaction(func(tx *query.Query) error {
		// 文件夹相关数据
		folderData := make([]rao.SaveTargetReq, 0, len(req.Apis))
		apiData := make([]rao.SaveTargetReq, 0, len(req.Apis))
		for _, info := range req.Apis {
			if info.TargetType == "folder" {
				folderData = append(folderData, info)
			}
			if info.TargetType == "api" {
				apiData = append(apiData, info)
			}
		}

		// 老的targetID对应新的targetID字典
		oldAndNewTargetIDMap := make(map[string]string)
		// 先把文件夹入库
		if len(folderData) > 0 {
			for _, folderInfo := range folderData {
				insertTargetInfo := packer.TransSaveImportFolderReqToTargetModel(folderInfo, req.TeamID, userID)
				newFolderName := insertTargetInfo.Name
				// 文件夹名排重
				_, err := tx.Target.WithContext(ctx).Where(tx.Target.TeamID.Eq(req.TeamID), tx.Target.Name.Eq(insertTargetInfo.Name),
					tx.Target.TargetType.Eq(consts.TargetTypeFolder), tx.Target.Status.Eq(consts.TargetStatusNormal)).First()
				if err == nil {
					newFolderName = newFolderName + "_1"
					// 查询老配置相关的
					list, err := tx.Target.WithContext(ctx).Where(tx.Target.TeamID.Eq(req.TeamID), tx.Target.Name.Like(fmt.Sprintf("%s%%", insertTargetInfo.Name+"_"))).Find()
					if err == nil && len(list) > 0 {
						// 有复制过得配置
						maxNum := 0
						for _, targetInfo := range list {
							nameTmp := targetInfo.Name
							postfixSlice := strings.Split(nameTmp, "_")
							if len(postfixSlice) < 2 {
								continue
							}
							currentNum, err := strconv.Atoi(postfixSlice[len(postfixSlice)-1])
							if err != nil {
								log.Logger.Info("复制自动化计划--类型转换失败，err:", err)
								continue
							}
							if currentNum > maxNum {
								maxNum = currentNum
							}
						}
						newFolderName = insertTargetInfo.Name + fmt.Sprintf("_%d", maxNum+1)
					}

				}
				// 插入target表
				insertTargetInfo.Name = newFolderName
				if folderInfo.OldParentID != "" {
					if newTargetID, ok := oldAndNewTargetIDMap[folderInfo.OldParentID]; ok {
						insertTargetInfo.ParentID = newTargetID
					}
				}

				err = tx.Target.WithContext(ctx).Create(insertTargetInfo)
				if err != nil {
					return err
				}
				oldAndNewTargetIDMap[folderInfo.OldTargetID] = insertTargetInfo.TargetID
			}
		}

		if len(apiData) > 0 {
			for _, apiInfo := range apiData {
				insertTargetInfo := packer.TransSaveImportTargetReqToTargetModel(apiInfo, req.TeamID, userID)
				newApiName := insertTargetInfo.Name
				// 接口名排重
				_, err := tx.Target.WithContext(ctx).Where(tx.Target.TeamID.Eq(req.TeamID), tx.Target.Name.Eq(insertTargetInfo.Name),
					tx.Target.TargetType.Eq(consts.TargetTypeAPI), tx.Target.Status.Eq(consts.TargetStatusNormal)).First()
				if err == nil {
					newApiName = newApiName + "_1"
					// 查询老配置相关的
					list, err := tx.Target.WithContext(ctx).Where(tx.Target.TeamID.Eq(req.TeamID), tx.Target.Name.Like(fmt.Sprintf("%s%%", insertTargetInfo.Name+"_"))).Find()
					if err == nil {
						// 有复制过得配置
						maxNum := 0
						for _, targetInfo := range list {
							nameTmp := targetInfo.Name
							postfixSlice := strings.Split(nameTmp, "_")
							if len(postfixSlice) < 2 {
								continue
							}
							currentNum, err := strconv.Atoi(postfixSlice[len(postfixSlice)-1])
							if err != nil {
								log.Logger.Info("复制自动化计划--类型转换失败，err:", err)
								continue
							}
							if currentNum > maxNum {
								maxNum = currentNum
							}
						}
						newApiName = insertTargetInfo.Name + fmt.Sprintf("_%d", maxNum+1)
					}

				}
				// 插入target表
				insertTargetInfo.Name = newApiName

				if apiInfo.OldParentID != "" {
					if newTargetID, ok := oldAndNewTargetIDMap[apiInfo.OldParentID]; ok {
						insertTargetInfo.ParentID = newTargetID
					}
				}

				err = tx.Target.WithContext(ctx).Create(insertTargetInfo)
				if err != nil {
					return err
				}

				apiDetail := packer.TransSaveTargetReqToMaoAPI(&apiInfo)
				// 把接口详情插入mongodb数据库
				apiDetail.TargetID = insertTargetInfo.TargetID
				_, err = collection.InsertOne(ctx, apiDetail)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	return err
}

func DetailByTargetIDs(ctx context.Context, teamID string, targetIDs []string) ([]*rao.APIDetail, error) {
	tx := query.Use(dal.DB()).Target
	targets, err := tx.WithContext(ctx).Where(
		tx.TargetID.In(targetIDs...),
		tx.TeamID.Eq(teamID),
		tx.TargetType.Eq(consts.TargetTypeAPI),
		tx.Status.Eq(consts.TargetStatusNormal),
		tx.Source.Eq(consts.TargetSourceNormal),
	).Order(tx.Sort.Desc(), tx.CreatedAt.Desc()).Find()

	if err != nil {
		return nil, err
	}

	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectAPI)
	cursor, err := collection.Find(ctx, bson.D{{"target_id", bson.D{{"$in", targetIDs}}}})

	if err != nil {
		return nil, err
	}
	var apis []*mao.API
	if err = cursor.All(ctx, &apis); err != nil {
		return nil, err
	}

	return packer.TransTargetsToRaoAPIDetails(targets, apis), nil
}

func CloneApi(ctx *gin.Context, req *rao.CloneApiReq) error {
	userID := jwt.GetUserIDByCtx(ctx)
	err := query.Use(dal.DB()).Transaction(func(tx *query.Query) error {
		oldApiInfo, err := tx.Target.WithContext(ctx).Where(tx.Target.TargetID.Eq(req.TargetID)).First()
		if err != nil {
			return err
		}

		oldApiName := oldApiInfo.Name   // 老接口名称
		newApiName := oldApiName + "_1" // 新接口名称

		// 查询老配置相关的
		targetNameList, err := tx.Target.WithContext(ctx).Where(tx.Target.TeamID.Eq(req.TeamID),
			tx.Target.TargetType.Eq(consts.TargetTypeAPI),
			tx.Target.Source.Eq(consts.TargetSourceNormal),
			tx.Target.Name.Like(fmt.Sprintf("%s%%", oldApiName+"_"))).Find()
		if err == nil && len(targetNameList) > 0 {
			// 有复制过得配置
			maxNum := 0
			for _, targetInfo := range targetNameList {
				nameTmp := targetInfo.Name
				postfixSlice := strings.Split(nameTmp, "_")
				if len(postfixSlice) < 2 {
					continue
				}
				currentNum, err := strconv.Atoi(postfixSlice[len(postfixSlice)-1])
				if err != nil {
					log.Logger.Info("复制接口--类型转换失败，err:", err)
					continue
				}
				if currentNum > maxNum {
					maxNum = currentNum
				}
			}
			newApiName = oldApiName + fmt.Sprintf("_%d", maxNum+1)
		}

		// 复制接口基本信息
		newTargetID := uuid.GetUUID()
		oldApiInfo.ID = 0
		oldApiInfo.TargetID = newTargetID
		oldApiInfo.Name = newApiName
		oldApiInfo.RecentUserID = userID
		oldApiInfo.CreatedUserID = userID
		oldApiInfo.CreatedAt = time.Now()
		oldApiInfo.UpdatedAt = time.Now()
		err = tx.Target.WithContext(ctx).Create(oldApiInfo)
		if err != nil {
			return err
		}

		// 复制接口详情信息
		var oldApiDetail mao.API
		collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectAPI)
		err = collection.FindOne(ctx, bson.D{{"target_id", req.TargetID}}).Decode(&oldApiDetail)
		if err == nil {
			oldApiDetail.TargetID = newTargetID
			_, err = collection.InsertOne(ctx, oldApiDetail)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
