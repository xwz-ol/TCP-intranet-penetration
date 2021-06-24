# server 与 client base module 与 function

首先我们先定义三个需要使用的工具方法，还需要定义两个消息编码常量，后面会用到

#######监听一个地址对应的 TCP 请求 CreateTCPListener
#######连接一个 TCP 地址 CreateTCPConn
#######将一个 TCP-A 连接的数据写入另一个 TCP-B 连接，将 TCP-B 连接返回的数据写入 TCP-A 的连接中 Join2Conn （别看这短短 10 几行代码，这就是核心了）
