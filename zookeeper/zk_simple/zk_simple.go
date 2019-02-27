package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var hosts = []string{"192.168.0.16:2181"}

func failToError(err error, hint string) {
	if err != nil {
		fmt.Println(hint, err)
	}
}
func ZkStateStringFormat(s *zk.Stat) string {
	return fmt.Sprintf("Czxid:%d\nMzxid: %d\nCtime: %d\nMtime: %d\nVersion: %d\nCversion: %d\nAversion: %d\nEphemeralOwner: %d\nDataLength: %d\nNumChildren: %d\nPzxid: %d\n",
		s.Czxid, s.Mzxid, s.Ctime, s.Mtime, s.Version, s.Cversion, s.Aversion, s.EphemeralOwner, s.DataLength, s.NumChildren, s.Pzxid)
}
func main() {
	fmt.Println("zookeeper start")
	conn, _, err := zk.Connect(hosts, time.Second*5)
	failToError(err, "connnect server false")
	defer conn.Close()

	//获取权限
	var acls = zk.WorldACL(zk.PermAll)

	var mypath = "/zk_mypath"
	var data = []byte("zk_mypath")
	var flag32 int32 = 0

	//创建目录
	_, create_err := conn.Create(mypath, data, flag32, acls)
	failToError(create_err, "create false")

	//获取路径下的数据
	path_data, stat, get_err := conn.Get(mypath)
	failToError(get_err, "get path error")

	fmt.Println(string(path_data))

	//判断是否存在
	b, stat, err := conn.Exists(mypath)
	if b {
		fmt.Println("路径存在")
	} else {
		fmt.Println("路径不存在")
	}
	//更新路径下的数据
	var new_data = []byte("new data")

	stat, err = conn.Set(mypath, new_data, stat.Version)
	failToError(err, "set mypath data false")
	ZkStateStringFormat(stat)
	//删除delete
	err = conn.Delete(mypath, stat.Version)
	failToError(err, "delete mypath")

}
