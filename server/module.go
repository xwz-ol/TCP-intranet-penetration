package server

import (
	"net"
	"sync"
	"time"
)

const (
	controlAddr = "0.0.0.0:8009"
	tunnelAddr  = "0.0.0.0:8008"
	visitAddr   = "0.0.0.0:8007"
)

var (
	clientConn         *net.TCPConn
	connectionPool     map[string]*ConnMatch
	connectionPoolLock sync.Mutex
)

type ConnMatch struct {
	addTime time.Time
	accept  *net.TCPConn
}
