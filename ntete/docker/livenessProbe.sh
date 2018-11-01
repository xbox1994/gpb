#!/bin/bash
# author: HuangChuanTong@WPS.cn
# date  : 2018-08-08
#
# 健康检测, 大概每10~15秒调用一次
#
#

HEALTHY_URL="http://127.0.0.1:8080/alive"

RESP=`curl --connect-timeout 1 -s ${HEALTHY_URL}`

if [ "$RESP" == "ok" ]; then
    echo "Server is alive."
    exit 0 # 服务正常，脚本退出: 0
else
    echo "Server is die."
    exit 1 # 脚本退出为非0,代表服务不可用
fi