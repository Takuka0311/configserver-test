package interfaces

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"

	configserverproto "configserver_test/proto"
)

func HeartBeat(r *gin.Engine, agent *configserverproto.Agent, requestID string) (int, *configserverproto.HeartBeatResponse) {
	// data
	reqBody := configserverproto.HeartBeatRequest{}
	reqBody.RequestId = requestID
	reqBody.AgentId = agent.AgentId
	reqBody.AgentType = agent.AgentType
	reqBody.RunningStatus = agent.RunningStatus
	reqBody.StartupTime = agent.StartupTime
	reqBody.Tags = agent.Tags
	reqBody.Attributes = agent.Attributes
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/Agent/HeartBeat", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.HeartBeatResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func FetchPipelineConfig(r *gin.Engine, configInfos []*configserverproto.ConfigCheckResult, requestID string) (int, *configserverproto.FetchPipelineConfigResponse) {
	// data
	reqBody := configserverproto.FetchPipelineConfigRequest{}
	reqBody.RequestId = requestID
	reqBody.ReqConfigs = make([]*configserverproto.ConfigInfo, 0)
	for _, c := range configInfos {
		conf := new(configserverproto.ConfigInfo)
		conf.Name = c.Name
		conf.Context = c.Context
		conf.Type = c.Type
		conf.Version = c.NewVersion
		reqBody.ReqConfigs = append(reqBody.ReqConfigs, conf)
	}
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/Agent/FetchPipelineConfig", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.FetchPipelineConfigResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}
