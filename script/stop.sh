#!/bin/bash

# 获取进程的PID
configserver_pid=$(ps aux | grep './ConfigServer' | grep -v grep | awk '{print $2}')
ilogtail_pid=$(ps aux | grep './ilogtail' | grep -v grep | awk '{print $2}')

# 杀死进程
if [ -n "$configserver_pid" ]; then
    kill $configserver_pid
    echo "configserver进程已杀死，PID为：$configserver_pid"
fi

if [ -n "$ilogtail_pid" ]; then
    kill $ilogtail_pid
    echo "ilogtail进程已杀死，PID为：$ilogtail_pid"
fi