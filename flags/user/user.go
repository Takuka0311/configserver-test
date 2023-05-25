package user

import (
	"configserver_test/interfaces"
	myproto "configserver_test/proto"
	"flag"
	"fmt"
	"os"
)

var (
	ConfigName  = flag.String("config-name", "", "")
	GroupName   = flag.String("group-name", "", "")
	Description = flag.String("description", "", "")
	Tags        = flag.String("tags", "", "")
	Context     = flag.String("context", "", "")
	Detail      = flag.String("detail", "", "")
)

func CreateAgentGroup() {

}

func UpdateAgentGroup() {

}

func DeleteAgentGroup() {
	funcName := "DeleteAgentGroup"
	name := *GroupName

	status, res := interfaces.DeleteAgentGroup(name, "test-"+funcName)

	printRes(funcName, status, res)
}

func GetAgentGroup() {
	funcName := "GetAgentGroup"
	name := *GroupName

	status, res := interfaces.GetAgentGroup(name, "test-"+funcName)

	printRes(funcName, status, res)
}

func ListAgentGroups() {

}

func CreateConfig() {
	funcName := "CreateConfig"
	name := *ConfigName
	detail := *Detail

	file, err := os.ReadFile(detail)
	fmt.Println(err)
	config := new(myproto.ConfigDetail)
	config.Name = name
	config.Type = myproto.ConfigType_PIPELINE_CONFIG
	config.Detail = string(file)

	status, res := interfaces.CreateConfig(config, "test-"+funcName)

	printRes(funcName, status, res)
}

func UpdateConfig() {
	funcName := "UpdateConfig"
	name := *ConfigName
	detail := *Detail

	file, err := os.ReadFile(detail)
	fmt.Println(err)
	config := new(myproto.ConfigDetail)
	config.Name = name
	config.Type = myproto.ConfigType_PIPELINE_CONFIG
	config.Detail = string(file)

	status, res := interfaces.UpdateConfig(config, "test-"+funcName)

	printRes(funcName, status, res)
}

func DeleteConfig() {
	funcName := "DeleteConfig"
	name := *ConfigName

	status, res := interfaces.DeleteConfig(name, "test-"+funcName)

	printRes(funcName, status, res)
}

func GetConfig() {
	funcName := "GetConfig"
	name := *ConfigName

	status, res := interfaces.GetConfig(name, "test-"+funcName)

	printRes(funcName, status, res)
}

func ListConfigs() {
	funcName := "ListConfigs"

	status, res := interfaces.ListConfigs("test-" + funcName)

	printRes(funcName, status, res)
}

func ApplyConfigToAgentGroup() {
	funcName := "ApplyConfigToAgentGroup"
	configName := *ConfigName
	groupName := *GroupName

	status, res := interfaces.ApplyConfigToAgentGroup(configName, groupName, "test-"+funcName)

	printRes(funcName, status, res)
}

func RemoveConfigFromAgentGroup() {
	funcName := "RemoveConfigFromAgentGroup"
	configName := *ConfigName
	groupName := *GroupName

	status, res := interfaces.RemoveConfigFromAgentGroup(configName, groupName, "test-"+funcName)

	printRes(funcName, status, res)
}

func GetAppliedConfigsForAgentGroup() {
	funcName := "GetAppliedConfigsForAgentGroup"
	name := *GroupName

	status, res := interfaces.GetAppliedConfigsForAgentGroup(name, "test-"+funcName)

	printRes(funcName, status, res)
}

func GetAppliedAgentGroups() {
	funcName := "GetAppliedAgentGroups"
	name := *ConfigName

	status, res := interfaces.GetAppliedAgentGroups(name, "test-"+funcName)

	printRes(funcName, status, res)
}

func ListAgents() {
	funcName := "ListAgents"
	name := *GroupName

	status, res := interfaces.ListAgents(name, "test-"+funcName)

	printRes(funcName, status, res)
}

func printRes(funcName string, status int, res interface{}) {
	fmt.Println(funcName)
	fmt.Println("-->status:", status)
	fmt.Println("-->response:", res)
	fmt.Println()
}
