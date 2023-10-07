package packer

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/log"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/mao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/model"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/go-omnibus/proof"
	"go.mongodb.org/mongo-driver/bson"
)

func TransSaveTargetReqToMaoAPI(target *rao.SaveTargetReq) *mao.API {
	if target.Request == nil {
		proof.Error("target.request not found request")
		return nil
	}

	header, err := bson.Marshal(target.Request.Header)
	if err != nil {
		proof.Error("target.request.header bson marshal err", proof.WithError(err))
	}

	query, err := bson.Marshal(target.Request.Query)
	if err != nil {
		log.Logger.Info("target.request.query bson marshal err", proof.WithError(err))
	}

	body, err := bson.Marshal(target.Request.Body)
	if err != nil {
		log.Logger.Info("target.request.body bson marshal err", proof.WithError(err))
	}

	auth, err := bson.Marshal(target.Request.Auth)
	if err != nil {
		log.Logger.Info("target.request.auth bson marshal err", proof.WithError(err))
	}

	assert, err := bson.Marshal(mao.Assert{Assert: target.Assert})
	if err != nil {
		log.Logger.Info("target.request.assert bson marshal err", proof.WithError(err))
	}

	regex, err := bson.Marshal(mao.Regex{Regex: target.Regex})
	if err != nil {
		log.Logger.Info("target.request.regex bson marshal err", proof.WithError(err))
	}

	return &mao.API{
		TargetID:      target.TargetID,
		PreUrl:        target.Request.PreUrl,
		URL:           target.Request.URL,
		EnvServiceID:  target.EnvServiceID,
		EnvServiceURL: target.EnvServiceURL,
		Header:        header,
		Query:         query,
		Body:          body,
		Auth:          auth,
		Description:   target.Description,
		Assert:        assert,
		Regex:         regex,
		HttpApiSetup: mao.HttpApiSetup{
			IsRedirects:  target.HttpApiSetup.IsRedirects,
			RedirectsNum: target.HttpApiSetup.RedirectsNum,
			ReadTimeOut:  target.HttpApiSetup.ReadTimeOut,
			WriteTimeOut: target.HttpApiSetup.WriteTimeOut,
		},
	}
}

func TransTargetToRaoAPIDetail(target *model.Target, api *mao.API, vs []*model.Variable) *rao.APIDetail {
	var auth rao.Auth
	if err := bson.Unmarshal(api.Auth, &auth); err != nil {
		log.Logger.Info("api.auth bson Unmarshal err", proof.WithError(err))
	}
	var body rao.Body
	if err := bson.Unmarshal(api.Body, &body); err != nil {
		log.Logger.Info("api.body bson Unmarshal err", proof.WithError(err))
	}
	var header rao.Header
	if err := bson.Unmarshal(api.Header, &header); err != nil {
		log.Logger.Info("api.header bson Unmarshal err", proof.WithError(err))
	}
	var query rao.Query
	if err := bson.Unmarshal(api.Query, &query); err != nil {
		log.Logger.Info("api.query bson Unmarshal err", proof.WithError(err))
	}

	var assert mao.Assert
	if err := bson.Unmarshal(api.Assert, &assert); err != nil {
		log.Logger.Info("api.assert bson Unmarshal err", proof.WithError(err))
	}

	var regex mao.Regex
	if err := bson.Unmarshal(api.Regex, &regex); err != nil {
		log.Logger.Info("api.regex bson Unmarshal err", proof.WithError(err))
	}

	var variables []*rao.KVVariable
	for _, v := range vs {
		variables = append(variables, &rao.KVVariable{
			Key:   v.Var,
			Value: v.Val,
		})
	}

	//// 检查是否有引用环境服务URL
	//var requestURLBuilder strings.Builder
	//if api.EnvServiceURL != "" {
	//	requestURLBuilder.WriteString(api.EnvServiceURL)
	//}
	//requestURLBuilder.WriteString(api.URL)
	//requestURL := requestURLBuilder.String()

	return &rao.APIDetail{
		TargetID:      target.TargetID,
		ParentID:      target.ParentID,
		TeamID:        target.TeamID,
		TargetType:    target.TargetType,
		Name:          target.Name,
		Method:        target.Method,
		URL:           api.URL,
		EnvServiceID:  api.EnvServiceID,
		EnvServiceURL: api.EnvServiceURL,
		Sort:          target.Sort,
		TypeSort:      target.TypeSort,
		Request: &rao.Request{
			PreUrl:      api.PreUrl,
			URL:         api.URL,
			Description: api.Description,
			Auth:        &auth,
			Body:        &body,
			Header:      &header,
			Query:       &query,
			Event:       nil,
			Cookie:      nil,
		},
		Response:       nil,
		Version:        target.Version,
		Description:    api.Description,
		CreatedTimeSec: target.CreatedAt.Unix(),
		UpdatedTimeSec: target.UpdatedAt.Unix(),
		Assert:         assert.Assert,
		Regex:          regex.Regex,
		Variable:       variables,
		HttpApiSetup: rao.HttpApiSetup{
			IsRedirects:  api.HttpApiSetup.IsRedirects,
			RedirectsNum: api.HttpApiSetup.RedirectsNum,
			ReadTimeOut:  api.HttpApiSetup.ReadTimeOut,
			WriteTimeOut: api.HttpApiSetup.WriteTimeOut,
		},
	}
}

func TransTargetsToRaoAPIDetails(targets []*model.Target, apis []*mao.API) []*rao.APIDetail {
	ret := make([]*rao.APIDetail, 0, len(targets))

	for _, target := range targets {
		for _, api := range apis {
			if api.TargetID == target.TargetID {
				ret = append(ret, TransTargetToRaoAPIDetail(target, api, nil))
			}
		}
	}

	return ret
}
