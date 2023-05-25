package main

import (
	"configserver_test/flags/agent"
	"configserver_test/flags/user"
	"flag"
	"fmt"
)

var Request = flag.String("request", "", "")

type Value interface {
	String() string
	Set(string) error
}

type sliceFlag []string

func (f *sliceFlag) String() string {
	return fmt.Sprintf("%v", []string(*f))
}

func (f *sliceFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}

var configs sliceFlag

func main() {
	flag.Var(&configs, "configs", "Pipline configs, for example: -configs=a.yaml -host=b.yaml")
	flag.Parse()

	switch *Request {
	case "CreateAgentGroup":
		user.CreateAgentGroup()
	case "UpdateAgentGroup":
		user.UpdateAgentGroup()
	case "DeleteAgentGroup":
		user.DeleteAgentGroup()
	case "GetAgentGroup":
		user.GetAgentGroup()
	case "ListAgentGroups":
		user.ListAgentGroups()
	case "CreateConfig":
		user.CreateConfig()
	case "UpdateConfig":
		user.UpdateConfig()
	case "DeleteConfig":
		user.DeleteConfig()
	case "GetConfig":
		user.GetConfig()
	case "ListConfigs":
		user.ListConfigs()
	case "ApplyConfigToAgentGroup":
		user.ApplyConfigToAgentGroup()
	case "RemoveConfigFromAgentGroup":
		user.RemoveConfigFromAgentGroup()
	case "GetAppliedConfigsForAgentGroup":
		user.GetAppliedConfigsForAgentGroup()
	case "GetAppliedAgentGroups":
		user.GetAppliedAgentGroups()
	case "ListAgents":
		user.ListAgents()
	case "HeartBeat":
		agent.HeartBeat(configs)
	}
}
