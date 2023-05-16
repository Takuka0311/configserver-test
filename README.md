# configserver-test使用帮助

## 启动 configserver和ilogtail

方便起见，直接放了一个编译好的可执行文件在configserver文件夹下，由于ilogtail超过了github文件大小限制，所以没有放，只在ilogtail文件夹下放了ilogtail_config.json。可以把用于测试的ilogtail复制进ilogtail目录下（ilogtil、libPluginAdapter.so、libPluginBase.h、libPluginBase.so）。

可以使用以下命令一键启动：

```shell
cd script
./start.sh 
# 如果显示Permission denied，需要使用如下命令加权限
# chmod 777 ./start.sh 
```

测试完成后记得关闭进程。

```shell
cd script
./stop.sh 
# 如果显示Permission denied，需要使用如下命令加权限
# chmod 777 ./stop.sh 
```

## 启动调试工具
