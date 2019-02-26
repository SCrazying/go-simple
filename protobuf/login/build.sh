#!/bin/bash
protoc --go_out=. login.proto
#安装库到环境中
go install 

protoc --go_out=. example.proto
#安装库到环境中
go install 