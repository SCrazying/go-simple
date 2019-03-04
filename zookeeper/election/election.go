package main

import (
	"errors"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"net"
	"strings"
	"time"
)

type ZookeeperConfig struct {
	Servers    []string
	RootPath   string
	MasterPath string
}
type ZookeeperManager struct {
	ZkClient *zk.Conn
	ZkConfig *ZookeeperConfig
	IsMaster chan bool
}

//初始化zookeeperManager实体
func newZookeeperManagerInstance(zkConfig *ZookeeperConfig, isMasterQ chan bool) *ZookeeperManager {
	manager := &ZookeeperManager{ZkClient: nil, ZkConfig: zkConfig, IsMaster: isMasterQ}
	manager.initConnect()
	return manager
}

//判断连接状态
func (zkManager *ZookeeperManager) isConnected() bool {

	if zkManager.ZkClient == nil {
		return false
	} else if zkManager.ZkClient.State() != zk.StateConnected {
		return false
	}
	return true
}

//选举master
func (zkManager *ZookeeperManager) electMaster() error {
	err := zkManager.initConnect()
	if err != nil {
		return err
	}
	isExists, _, err := zkManager.ZkClient.Exists(zkManager.ZkConfig.RootPath)
	if err != nil {
		return err
	}
	if !isExists {
		path, err := zkManager.ZkClient.Create(zkManager.ZkConfig.RootPath, nil, 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			return nil
		}
		if zkManager.ZkConfig.RootPath != path {
			return errors.New("Create returned different path " + zkManager.ZkConfig.RootPath + " != " + path)
		}
	}
	//创建选取用的master的ZNode节点，节点类型是Ephemeral，这种类型是断开连接后，创建的节点对应被销毁
	masterpath := zkManager.ZkConfig.RootPath + zkManager.ZkConfig.MasterPath
	path, err := zkManager.ZkClient.Create(masterpath, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		fmt.Printf("elect master failure, ", err)
		zkManager.IsMaster <- false
		return err
	} else {
		if path == masterpath {
			fmt.Println("create zknode success")
			zkManager.IsMaster <- true
		} else {
			return errors.New("Create returned defferent path " + masterpath + " != " + path)
		}
	}
	return nil
}

//
func (zkManager *ZookeeperManager) run() {
	//开始选举master
	err := zkManager.electMaster()
	if err != nil {
		fmt.Println("elect master error", err)
	}
	zkManager.watchMaster()
}

//初始化连接
func (zkManager *ZookeeperManager) initConnect() error {

	if !zkManager.isConnected() {

		//连接
		conn, connChan, err := zk.Connect(zkManager.ZkConfig.Servers, time.Second)
		if err != nil {
			log.Fatalln("connect hosts false ")

		}
		for {
			isConnected := false
			select {
			case connEvent := <-connChan:
				if connEvent.State == zk.StateConnected {
					isConnected = true
					fmt.Println("connect hosts success")
				}
			case _ = <-time.After(time.Second * 3):
				fmt.Println("timeout 3s")
				return errors.New("timeout 3s")
			}
			if isConnected {
				break
			}

		}
		zkManager.ZkClient = conn
	}
	return nil
}

func (zkManager *ZookeeperManager) watchMaster() error {
	masterpath := zkManager.ZkConfig.RootPath + zkManager.ZkConfig.MasterPath
	children, stat, connChan, err := zkManager.ZkClient.ChildrenW(masterpath)

	if err != nil {
		return err
	}
	fmt.Println("watch children result, ", children, stat)
	for {
		select {
		case evt := <-connChan:
			//节点退出
			if evt.Type == zk.EventNodeDeleted {

				fmt.Println("receive znode delete event, ", evt)
				// 重新选举
				fmt.Println("start elect new master ...")
				err = zkManager.electMaster()
				if err != nil {
					fmt.Println("elect new master error, ", err)
				}
			}
		}
	}
	return nil
}

func localHostIp() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get ip failed")
		return nil, nil
	}
	data := make([]string, 0)
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				data = append(data, ipnet.String())
			}
		}
	}
	return data, nil
}
func main() {

	//zkConfig
	zkConfig := &ZookeeperConfig{
		Servers:    []string{"192.168.0.16:2181"},
		RootPath:   "/ElectMaster",
		MasterPath: "/master",
	}

	isMasterChan := make(chan bool)

	zkManager := newZookeeperManagerInstance(zkConfig, isMasterChan)
	go zkManager.run()
	var isMaster bool
	ips, _ := localHostIp()
	for {
		select {
		case isMaster = <-isMasterChan:
			if isMaster {
				fmt.Println("hosts" + strings.Join(ips, " ") + " is master")
			}
		}
	}
}
