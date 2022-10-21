package response

import (
	"net"
	"time"
)

type Trace struct {
	DNSLookup      time.Duration
	ConnTime       time.Duration
	TCPConnTime    time.Duration
	TLSHandshake   time.Duration
	ServerTime     time.Duration
	ResponseTime   time.Duration
	TotalTime      time.Duration
	IsConnReused   bool
	IsConnWasIdle  bool
	ConnIdleTime   time.Duration
	RequestAttempt int
	RemoteAddr     net.Addr
}
