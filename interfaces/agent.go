package interfaces

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/protobuf/proto"

	configserverproto "configserver_test/proto"
)

func HeartBeat(agent *configserverproto.Agent, agentConfigs []*configserverproto.ConfigInfo,
	pipelineConfigs []*configserverproto.ConfigInfo, requestID string) (int, *configserverproto.HeartBeatResponse) {
	// data
	reqBody := configserverproto.HeartBeatRequest{}
	reqBody.RequestId = requestID
	reqBody.AgentConfigs = agentConfigs
	reqBody.AgentId = agent.GetAgentId()
	reqBody.AgentType = agent.GetAgentType()
	reqBody.Attributes = agent.GetAttributes()
	reqBody.Interval = agent.GetInterval()
	reqBody.PipelineConfigs = pipelineConfigs
	reqBody.RunningStatus = agent.GetRunningStatus()
	reqBody.StartupTime = agent.GetStartupTime()
	reqBody.Tags = agent.GetTags()
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8899/Agent/HeartBeat", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.HeartBeatResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	fmt.Println(res.Status)

	return res.StatusCode, resBody
}

func FetchPipelineConfig(configInfos []*configserverproto.ConfigCheckResult, requestID string) (int, *configserverproto.FetchPipelineConfigResponse) {
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
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8899/Agent/FetchPipelineConfig", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.FetchPipelineConfigResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}
