package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"log"
	"time"
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

	fmt.Println("go client start")
	//使用点对点通信
	d := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	//创建一个XClient
	xclient := client.NewXClient("Arith",
		client.Failtry,       //失败重试
		client.RandomSelect,  //服务调用算法
		d,                    //服务调用方式，点对点，一对多，concul，etcd，zookeeper
		client.DefaultOption, //客户端配置
	)

	args := Args{
		A: 10,
		B: 30,
	}
	//定义返回的结构
	reply := &Reply{}
	//异步调用rpc接口，返回一个chan用于获取数据
	replyCh, err := xclient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Println("call error")
	}
	//等待数据返回或者超时退出
	for {
		var flag = false
		select {
		case evt := <-replyCh.Done:
			if evt.Error != nil {
				log.Fatalf("failed to call: %v", evt.Error)
			} else {
				fmt.Printf("%d * %d = %d\n", args.A, args.B, reply.C)
			}
			flag = true
		case <-time.After(time.Second):
			log.Fatalf("timeout to call")
			flag = true
		}
		if flag {
			break
		}
	}

}
