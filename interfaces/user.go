package interfaces

import (
	"bytes"
	"io"
	"net/http"

	"google.golang.org/protobuf/proto"

	configserverproto "configserver_test/proto"
)

func DeleteAgentGroup(groupName string, requestID string) (int, *configserverproto.DeleteAgentGroupResponse) {
	// data
	reqBody := configserverproto.DeleteAgentGroupRequest{}
	reqBody.RequestId = requestID
	reqBody.GroupName = groupName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", "http://127.0.0.1:8899/User/DeleteAgentGroup", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.DeleteAgentGroupResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func GetAgentGroup(groupName string, requestID string) (int, *configserverproto.GetAgentGroupResponse) {
	// data
	reqBody := configserverproto.GetAgentGroupRequest{}
	reqBody.RequestId = requestID
	reqBody.GroupName = groupName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8899/User/GetAgentGroup", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.GetAgentGroupResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func CreateConfig(config *configserverproto.ConfigDetail, requestID string) (int, *configserverproto.CreateConfigResponse) {
	// data
	reqBody := configserverproto.CreateConfigRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigDetail = config
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8899/User/CreateConfig", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.CreateConfigResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func GetConfig(configName string, requestID string) (int, *configserverproto.GetConfigResponse) {
	// data
	reqBody := configserverproto.GetConfigRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigName = configName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8899/User/GetConfig", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response
	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.GetConfigResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func UpdateConfig(config *configserverproto.ConfigDetail, requestID string) (int, *configserverproto.UpdateConfigResponse) {
	// data
	reqBody := configserverproto.UpdateConfigRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigDetail = config
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", "http://127.0.0.1:8899/User/UpdateConfig", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response

	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.UpdateConfigResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func DeleteConfig(configName string, requestID string) (int, *configserverproto.DeleteConfigResponse) {
	// data
	reqBody := configserverproto.DeleteConfigRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigName = configName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", "http://127.0.0.1:8899/User/DeleteConfig", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response

	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.DeleteConfigResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func ListConfigs(requestID string) (int, *configserverproto.ListConfigsResponse) {
	// data
	reqBody := configserverproto.ListConfigsRequest{}
	reqBody.RequestId = requestID
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8899/User/ListConfigs", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response

	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.ListConfigsResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func ApplyConfigToAgentGroup(configName string, groupName string, requestID string) (int, *configserverproto.ApplyConfigToAgentGroupResponse) {
	// data
	reqBody := configserverproto.ApplyConfigToAgentGroupRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigName = configName
	reqBody.GroupName = groupName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", "http://127.0.0.1:8899/User/ApplyConfigToAgentGroup", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response

	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.ApplyConfigToAgentGroupResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func RemoveConfigFromAgentGroup(configName string, groupName string, requestID string) (int, *configserverproto.RemoveConfigFromAgentGroupResponse) {
	// data
	reqBody := configserverproto.RemoveConfigFromAgentGroupRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigName = configName
	reqBody.GroupName = groupName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", "http://127.0.0.1:8899/User/RemoveConfigFromAgentGroup", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response

	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.RemoveConfigFromAgentGroupResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func GetAppliedConfigsForAgentGroup(groupName string, requestID string) (int, *configserverproto.GetAppliedConfigsForAgentGroupResponse) {
	// data
	reqBody := configserverproto.GetAppliedConfigsForAgentGroupRequest{}
	reqBody.RequestId = requestID
	reqBody.GroupName = groupName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8899/User/GetAppliedConfigsForAgentGroup", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response

	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.GetAppliedConfigsForAgentGroupResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func GetAppliedAgentGroups(configName string, requestID string) (int, *configserverproto.GetAppliedAgentGroupsResponse) {
	// data
	reqBody := configserverproto.GetAppliedAgentGroupsRequest{}
	reqBody.RequestId = requestID
	reqBody.ConfigName = configName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8899/User/GetAppliedAgentGroups", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response

	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.GetAppliedAgentGroupsResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}

func ListAgents(groupName string, requestID string) (int, *configserverproto.ListAgentsResponse) {
	// data
	reqBody := configserverproto.ListAgentsRequest{}
	reqBody.RequestId = requestID
	reqBody.GroupName = groupName
	reqBodyByte, _ := proto.Marshal(&reqBody)

	// request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8899/User/ListAgents", bytes.NewBuffer(reqBodyByte))
	res, _ := client.Do(req)

	// response

	resBodyByte, _ := io.ReadAll(res.Body)
	resBody := new(configserverproto.ListAgentsResponse)
	_ = proto.Unmarshal(resBodyByte, resBody)

	return res.StatusCode, resBody
}
