package client

var (

	// 本地暴露服务地址
	localServerAddr = "127.0.0.1:32768"
	// 远端服务控制通道， 用于传递控制信息，如出现新连接和心跳
	remoteIP          = "111.111.111.111"
	remoteControlAddr = remoteIP + ":8008"

	// 远端服务端口， 用来建立隧道
	remoteServerAddr = remoteIP + "8009"
)
