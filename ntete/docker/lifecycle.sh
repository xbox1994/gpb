#!/bin/bash
# author: HuangChuanTong@WPS.cn
# date  : 2018-08-08
#
# 退出回收, 发送kill -TERM ${pid} 后10~15秒调用脚本检测
# SIGTERM=15
#

HEALTHY_URL="http://127.0.0.1:8080/alive"

RESP=`curl --connect-timeout 3 -s ${HEALTHY_URL}`

if [ "$RESP" == "ok" ]; then
    echo "Server is running."
    exit 1 # 服务还在运行，脚本退出: 非0，
else
    echo "Server was exit."
    exit 0 # 服务已关闭并退出，脚本退出：0
           # 接着会回收容器资源。
fi
