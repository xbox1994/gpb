#!/bin/bash
# author: HuangChuanTong@WPS.cn
# date  : 2018-08-08
#
# 就绪检测, 大概start.sh执行后10~15秒调用，
#
#

HEALTHY_URL="http://127.0.0.1:8080/alive"

RESP=`curl --connect-timeout 1 -s ${HEALTHY_URL}`

if [ "$RESP" == "ok" ]; then
    echo "Server is ready."
    exit 0 # 服务已准备好，脚本退出: 0，
            # 接着会放流量，不再调用此脚本。

else
    echo "Server is not ready."
    exit 1 # 脚本退出为非0,代表服务还不可用，
           # 不会放流量进来，前重复调用脚本
fi