#!/bin/bash

#----------------------------------------
# 构建机仅内置了有限的go版本，用最新版本，联系SRE团队。
GO_VERSION='1.10.3'
PROJECT_NAME="ntete"

#  构建机会预先拉取工程代码，如工程叫  sample1
#  git clone git@ksogit.kingsoft.net:pub/sample1.git sample1
#  cd ./sample1
#  以下指令起始目录是 工程根目录。

#----------------------------------------
# 导出 PROJECT_ROOT 请不要更改此段shell
if [ ! -e "$PWD" ] || [ "$PWD" == "" ] || [ "$PWD" == "/" ] ; then
    echo "PROJECT_ROOT env did not set, It's empty string. exit -1"
    exit -1
else
   export PROJECT_ROOT="$PWD"
fi
cd "$PROJECT_ROOT"



# 构建机上，不同版本golang安装目录固定在： /usr/local/
# 如 /usr/local/go1.10.3 ; 即可导出GOROOT与PATH
export GOROOT=/usr/local/go${GO_VERSION}/go
export GOTOOLDIR="$GOROOT/pkg/tool/linux_amd64"
export PATH=$GOROOT/bin:$PATH
export GOPATH=${PROJECT_ROOT}/gopath
export GOBIN=${PROJECT_ROOT}/bin

#clear workspace
rm -rf ${PROJECT_ROOT}/gopath
rm -rf ${PROJECT_ROOT}/bin

if [ ! -h ${GOPATH}/src/${PROJECT_NAME} ]; then
    mkdir -p ${GOPATH}/src
    ln -s "${PROJECT_ROOT}" ${GOPATH}/src/${PROJECT_NAME}
fi

#公共基础库
git clone git@ksogit.kingsoft.net:mo_server/vendor.git ${GOPATH}/src/vendor

cd ${GOPATH}/src/${PROJECT_NAME}

go install .