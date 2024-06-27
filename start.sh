#! /bin/sh

CONFIG_NAME=$1
ERRMSG="config.yml|config.test.yml"
if [ ${ENV_NAME} = "" ]; then
    echo $ERRMSG
    exit
fi

echo "================start shell==================="

killall -9 merchant

nohup ./merchant -c=${CONFIG_NAME} > Nohup.log &

ps -ef | grep merchant | grep  -v grep

echo "================end shell==================="