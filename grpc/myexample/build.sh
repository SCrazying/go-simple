#!/bin/bash

protoc --go_out=plugins=grpc:. example.proto
#安装库到环境中
go install 