package agent

import (
	"configserver_test/interfaces"
	myproto "configserver_test/proto"
	"flag"
	"fmt"
)

var (
	AgentId = flag.String("agent-id", "", "")
)

func HeartBeat(configs []string) {
	funcName := "HeartBeat"
	id := *AgentId
	pipelinConfigs := make([]*myproto.ConfigInfo, len(configs))
	for i := 0; i < len(configs); i++ {
		pipelinConfigs[i] = new(myproto.ConfigInfo)
		pipelinConfigs[i].Name = configs[i]
		pipelinConfigs[i].Type = myproto.ConfigType_PIPELINE_CONFIG
		pipelinConfigs[i].Version = 0
		pipelinConfigs[i].Context = ""
	}

	attrs := new(myproto.AgentAttributes)
	attrs.Category = ""
	attrs.Extras = make(map[string]string, 0)
	attrs.Hostname = ""
	attrs.Ip = ""
	attrs.Region = ""
	attrs.Version = "0.0"
	attrs.Zone = ""

	agent := new(myproto.Agent)
	agent.AgentId = id
	agent.AgentType = "testAgent"
	agent.Attributes = attrs
	agent.Interval = 10
	agent.RunningStatus = ""
	agent.StartupTime = 0
	agent.Tags = make([]string, 0)

	status, res := interfaces.HeartBeat(agent, make([]*myproto.ConfigInfo, 0), pipelinConfigs, "test-"+funcName)

	fmt.Println(funcName)
	fmt.Println("-->status:", status)
	fmt.Println("-->response:", res)
	if len(res.PipelineCheckResults) != 0 {
		for i, result := range res.PipelineCheckResults {
			fmt.Println("-->PipelineCheckResults[", i, "]: Name=", result.Name, ", CheckStatus=", result.CheckStatus,
				", OldVersion=", result.OldVersion, ", NewVersion=", result.NewVersion)
		}
		FetchPipelineConfig(res.PipelineCheckResults)
	}
	fmt.Println()
}

func FetchPipelineConfig(PipelineCheckResults []*myproto.ConfigCheckResult) {
	funcName := "FetchPipelineConfig"

	status, res := interfaces.FetchPipelineConfig(PipelineCheckResults, "test-"+funcName)

	printRes(funcName, status, res)
}

func printRes(funcName string, status int, res interface{}) {
	fmt.Println(funcName)
	fmt.Println("-->status:", status)
	fmt.Println("-->response:", res)
	fmt.Println()
}
