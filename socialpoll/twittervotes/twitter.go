package twitter

import (
	"net"
	"time"
)

var conn net.Conn

func dial(netw, addr string) (net.Conn, error) {
	if conn != nil { //close the connection and set it to nil if it's not closed properly before
		conn.Close()
		conn = nil
	}

	netc, err := net.DialTimeout(netw, addr, 3*time.Second)

	if err != nil { //if error then return nil connection and error
		return nil, err
	}

	conn = netc
	return netc, nil
}
