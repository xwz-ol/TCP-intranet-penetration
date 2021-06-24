package client

import (
	"TCP-intranet-penetration/base"
	"bufio"
	"fmt"
	"net"
)

func NewClient() {
	client()
	return
}

func client() {
	tcpConn, err := base.CreateTcpConn(remoteControlAddr)
	if err != nil {
		fmt.Println("client 连接 %v err: %v", remoteControlAddr, err)
		return
	}

	reader := bufio.NewReader(tcpConn)

	for {
		content, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("reader ReadString err: %v", err)
			break
		}

		if content == base.NewConnection+"\n" {
			// todo 新建一个tcp conn
			go createConnect()
		}

	}

	fmt.Println("connect is broken" + remoteServerAddr)
}

// createConnect
func createConnect() {
	sourceConn := createConnectBase(localServerAddr)
	targetConn := createConnectBase(remoteServerAddr)
	// 数据交换
	if sourceConn == nil || targetConn == nil {
		_ = sourceConn.Close()
		_ = targetConn.Close()
		fmt.Println("createConnect is failed")
		return
	}
	base.CopyConn(sourceConn, targetConn)

}

// createConnectBase
func createConnectBase(addr string) *net.TCPConn {
	conn, err := base.CreateTcpConn(addr)
	if err != nil {
		fmt.Println("connect building is err: %v", err)
	}

	return conn
}
