package base

import (
	"fmt"
	"io"
	"net"
)

const (
	KeepAlive     = "KEEP_ALIVE"
	NewConnection = "NEW_CONNECTION"
)

// 根据传入受监听地址端口返回一个tcp Listener
func CreateTcpListenner(addr string) (*net.TCPListener, error) {

	// 校验IPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Println("CreateTcoListenner ResolveIPAddr tcpAddr err %v", err)
		panic(err)
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("CreateTcoListenner ListenTCP is err %v", err)
		return nil, err
	}

	return tcpListener, err
}

// 根据目标地址创建一个TCP Conn
func CreateTcpConn(target string) (*net.TCPConn, error) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		fmt.Println("CreateTcpConn ResolveTCPAddr target err %v", err)
		panic(err)
	}
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("CreateTcpConn DialTCP is err %v", err)
		return nil, err
	}

	return tcpConn, err
}

// 对外暴露异步数据传输函数
func CopyConn(sourceConn *net.TCPConn, targetConn *net.TCPConn) {
	go copyConn(sourceConn, targetConn)
	go copyConn(targetConn, sourceConn)
}

// 数据传输
func copyConn(sourceConn *net.TCPConn, targetConn *net.TCPConn) {
	defer sourceConn.Close()
	defer targetConn.Close()

	_, err := io.Copy(sourceConn, targetConn)
	if err != nil {
		fmt.Println("CreateTcpConn DialTCP is err %v", err)
		return
	}

}
