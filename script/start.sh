#!/bin/bash

# start config server
cd ../configserver
nohup ./ConfigServer > stdout.log 2> stderr.log &
config_pid=$!

# start ilogtail
cd ../ilogtail
nohup ./ilogtail > stdout.log 2> stderr.log &
ilog_pid=$!

# output PIDs
echo "Config server PID: $config_pid"
echo "Ilogtail PID: $ilog_pid"
