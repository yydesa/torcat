package main

import (
	"conntools"
	"golang.org/x/net/proxy"
	"net"
	"os"
)

// I suppose there is something easier for this wrappering,
// but I don't know it.

type MD struct {
}

func (m MD) Dial(network, addr string) (c net.Conn, err error) {
	return net.Dial(network, addr)
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("No destination")
	}
	host := args[0]
	d, err := proxy.SOCKS5("tcp", "127.0.0.1:9150", nil, MD{})
	if err != nil {
		panic("failed to make dialer: " + err.Error())
	}

	conn, err := d.Dial("tcp", host)
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	conntools.Run(conn)
}
