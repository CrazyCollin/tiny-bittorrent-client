package client

import "net"

type Client struct {
	conn     net.Conn
	infoHash [20]byte
	peerID   [20]byte
}
