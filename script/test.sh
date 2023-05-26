cd ..
# go build
# ./configserver_test --request=CreateConfig --config-name=test.yaml --detail=./test_configs/test.yaml
# ./configserver_test --request=UpdateConfig --config-name=test.yaml --detail=./test_configs/test.yaml
# ./configserver_test --request=DeleteConfig --config-name=test.yaml
# ./configserver_test --request=DeleteAgentGroup --group-name=default
# ./configserver_test --request=ApplyConfigToAgentGroup --config-name=test.yaml --group-name=default
./configserver_test --request=RemoveConfigFromAgentGroup --config-name=test.yaml --group-name=default
# ./configserver_test --request=HeartBeat --agent-id=test-1
./configserver_test --request=HeartBeat --agent-id=test-1 -configs=test.yaml 
# ./configserver_test --request=GetAgentGroup --group-name=default
# ./configserver_test --request=ListAgents --group-name=default
# ./configserver_test --request=ListConfigs
# ./configserver_test --request=GetConfig --config-name=test.yaml
./configserver_test --request=GetAppliedConfigsForAgentGroup --group-name=default