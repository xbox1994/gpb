#!/bin/bash
# author: HuangChuanTong@WPS.cn
# date  : 2018-08-08
#
# docker 容器启动时调用的第一个脚本，自行定制要启动的服务。
#   PROJECT 为平台内置环境变量，
#   supervisord 默认安装，用于服务管理。

mkdir /dev/shm/logs -p

alias cp="cp"
cp -rf /opt/apps/$PROJECT/docker/supervisord/apps.conf  /etc/supervisord/

# 启动supervisor 管理服务
/usr/bin/supervisord
