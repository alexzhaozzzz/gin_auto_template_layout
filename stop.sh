#! /bin/sh

echo "================start shell==================="

killall -2 merchant

ps -ef | grep merchant | grep  -v grep

echo "================end shell==================="