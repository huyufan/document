#!/bin/sh

ps aux | grep testcsgo

count=$(ps -ef | grep "testcsgo" | grep -v "grep" | wc -l)

echo ""

if [ $count -eq 0 ];then
    echo "testcsgo starting..."
    ./testcsgo &
    echo "testcsgo started"
else
    echo "testcsgo Restarting..."
    kill -USR2 $(ps -ef | grep "testcsgo" | grep -v grep | awk '{print $2}')
    echo "testcsgo Restarted"
fi

sleep 1

ps aux | grep testcsgo
