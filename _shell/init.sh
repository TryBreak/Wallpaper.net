#!/bin/bash

# 项目根目录
path=$(pwd)

# 项目的名字和编译时的名字
startName=${path##*/}
buildName=${startName}

# log 目录
logPath=${path}"/logs"

# 最终的输出目录
outPutPath=${path}"/dist"

# 配置文件
userEnv=${path}"/user_config.yaml"

# 部署目录
deployPath="/root/ProdProject/"${startName}
