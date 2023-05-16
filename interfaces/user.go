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

func DeleteAgentGroup(r *gin.Engine, groupName string, requestID string) (int, *configserverproto.DeleteAgentGroupResponse) {
	// data
	reqBody := configserverproto.DeleteAgentGroupRequest{}
	reqBody.RequestId = requestID
	reqBody.GroupName = groupName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/User/DeleteAgentGroup", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.DeleteAgentGroupResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func CreateConfig(r *gin.Engine, config *configserverproto.ConfigDetail, requestID string) (int, *configserverproto.CreateConfigResponse) {
	// data
	reqBody := configserverproto.CreateConfigRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigDetail = config
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/User/CreateConfig", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.CreateConfigResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func GetConfig(r *gin.Engine, configName string, requestID string) (int, *configserverproto.GetConfigResponse) {
	// data
	reqBody := configserverproto.GetConfigRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigName = configName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/User/GetConfig", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.GetConfigResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func UpdateConfig(r *gin.Engine, config *configserverproto.ConfigDetail, requestID string) (int, *configserverproto.UpdateConfigResponse) {
	// data
	reqBody := configserverproto.UpdateConfigRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigDetail = config
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/User/UpdateConfig", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.UpdateConfigResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func DeleteConfig(r *gin.Engine, configName string, requestID string) (int, *configserverproto.DeleteConfigResponse) {
	// data
	reqBody := configserverproto.DeleteConfigRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigName = configName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/User/DeleteConfig", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.DeleteConfigResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func ListConfigs(r *gin.Engine, requestID string) (int, *configserverproto.ListConfigsResponse) {
	// data
	reqBody := configserverproto.ListConfigsRequest{}
	reqBody.RequestId = requestID
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/User/ListConfigs", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.ListConfigsResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func ApplyConfigToAgentGroup(r *gin.Engine, configName string, groupName string, requestID string) (int, *configserverproto.ApplyConfigToAgentGroupResponse) {
	// data
	reqBody := configserverproto.ApplyConfigToAgentGroupRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigName = configName
	reqBody.GroupName = groupName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/User/ApplyConfigToAgentGroup", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.ApplyConfigToAgentGroupResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func RemoveConfigFromAgentGroup(r *gin.Engine, configName string, groupName string, requestID string) (int, *configserverproto.RemoveConfigFromAgentGroupResponse) {
	// data
	reqBody := configserverproto.RemoveConfigFromAgentGroupRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigName = configName
	reqBody.GroupName = groupName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/User/RemoveConfigFromAgentGroup", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.RemoveConfigFromAgentGroupResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func GetAppliedConfigsForAgentGroup(r *gin.Engine, groupName string, requestID string) (int, *configserverproto.GetAppliedConfigsForAgentGroupResponse) {
	// data
	reqBody := configserverproto.GetAppliedConfigsForAgentGroupRequest{}
	reqBody.RequestId = requestID
	reqBody.GroupName = groupName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/User/GetAppliedConfigsForAgentGroup", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.GetAppliedConfigsForAgentGroupResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func GetAppliedAgentGroups(r *gin.Engine, configName string, requestID string) (int, *configserverproto.GetAppliedAgentGroupsResponse) {
	// data
	reqBody := configserverproto.GetAppliedAgentGroupsRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigName = configName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/User/GetAppliedAgentGroups", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.GetAppliedAgentGroupsResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func ListAgents(r *gin.Engine, groupName string, requestID string) (int, *configserverproto.ListAgentsResponse) {
	// data
	reqBody := configserverproto.ListAgentsRequest{}
	reqBody.RequestId = requestID
	reqBody.GroupName = groupName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/User/ListAgents", bytes.NewBuffer(reqBodyByte))
	r.ServeHTTP(w, req)

	// response
	res := w.Result()
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.ListAgentsResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}
