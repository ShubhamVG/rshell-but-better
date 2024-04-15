package server

import "net"

// This is just localAddr + remoteAddr
type UniqueConnAddr = string

// I hope this is unique enuff
func getUniqueConnAddr(connPtr *net.Conn) UniqueConnAddr {
	conn := *connPtr

	return conn.LocalAddr().String() + " " + conn.RemoteAddr().String()
}
