# configserver-test 使用帮助

## 启动 configserver

方便起见，直接放了一个编译好的可执行文件在configserver文件夹下。

### 仅启动 configserver

```shell
cd configserver
nohup ./ConfigServer > stdout.log 2> stderr.log &
```

测试完后记得把进程干掉。

### 同时启动 configserver 和 ilogtail

如果想与ilogtail连调，可以采用此方法。由于ilogtail超过了github文件大小限制，所以没有上传编译好的可执行文件，只在ilogtail文件夹下放了ilogtail_config.json。可以把用于测试的ilogtail复制进ilogtail目录下（ilogtil、libPluginAdapter.so、libPluginBase.h、libPluginBase.so）。注：需要使用最新的[开源ilogtail代码](https://github.com/alibaba/ilogtail)编译。

可以使用以下命令一键启动：

```shell
cd script
./start.sh 
# 如果显示Permission denied，需要使用如下命令加权限，下面同理
# chmod 777 ./start.sh 
```

可以用clean脚本清除产生的数据

```shell
cd script
./clean.sh 
```

测试完成后记得关闭进程。

```shell
cd script
./stop.sh 
```

## 使用调试工具

编译工具

```shell
go build
```

在启动ConfigServer和logtail后，使用命令模拟发送API请求。例如：

```shell
./configserver_test --request=CreateConfig --config-name=test.yaml --detail=./test_configs/test.yaml
```

也可以参考script文件夹下的test.sh文件，将多个命令批量进行。

```shell
cd script
./test.sh 
```
