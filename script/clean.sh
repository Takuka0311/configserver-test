#!/bin/bash

# 清理configserver文件夹
rm -rf ../configserver/DB
find ../configserver/ -name "*.log" -type f -delete

# 清理ilogtail文件夹
rm -rf ../ilogtail/checkpoint
rm -rf ../ilogtail/snapshot
rm -rf ../ilogtail/user_config.d
rm -rf ../ilogtail/user_yaml_config.d
rm -rf ../ilogtail/users
rm -rf ../ilogtail/app_info.json
rm -rf ../ilogtail/apsara_log_conf.json
rm -rf ../ilogtail/backtrace.dat
find ../ilogtail/ -name "*.log" -type f -delete
find ../ilogtail/ -name "*.LOG" -type f -delete