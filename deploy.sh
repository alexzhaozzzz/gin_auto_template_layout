#!/bin/bash
# author Chris
cn="br"
PROJECT_NAME=$(basename "$(pwd)")
file_path="/usr/local/gosrc/src/back-end"
end_file_path="/usr/local/slots/release"
SERVER='10.1.22.193'
echo "当前项目名称：$PROJECT_NAME"

md5sum ${file_path}/${PROJECT_NAME}/release/${PROJECT_NAME}.bin 
scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -r -p -o LogLevel=ERROR -o ConnectTimeout=120 -oServerAliveInterval=60 -o ServerAliveCountMax=10 -B -C -q  -o TCPKeepAlive=yes   ${file_path}/${PROJECT_NAME}/release/${PROJECT_NAME}.bin root@${SERVER}:${end_file_path}/${PROJECT_NAME}/br-${PROJECT_NAME}.bin 
ssh root@${SERVER} "sudo bash /opt/script/project-deploy.sh ${cn} ${PROJECT_NAME}"