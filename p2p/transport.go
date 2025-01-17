package p2p

import "net"

type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
	Dial(string) error
	ListenAddr() string
	Close() error
}
