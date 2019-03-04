package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"log"
)

const addr = "127.0.0.1:2333"

type Args struct {
	A int
	B int
}
type Reply struct {
	C int
}

func main() {

	//使用点对点通信
	d := client.NewPeer2PeerDiscovery("tcp@"+addr, "")

	//创建一个XClient
	xclient := client.NewXClient("Arith",
		client.Failtry,       //失败重试
		client.RandomSelect,  //服务调用算法
		d,                    //服务调用方式，点对点，一对多，concul，etcd，zookeeper
		client.DefaultOption, //客户端配置
	)
	//定义请求体
	args := Args{
		A: 10,
		B: 30,
	}
	//定义返回的结构
	reply := &Reply{}
	//调用rpc接口
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Println("call error")
	}
	fmt.Printf("%d * %d = %d\n", args.A, args.B, reply.C)

}
