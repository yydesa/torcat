package main

import (
	"conntools"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("No destination")
	}
	host := args[0]

	conn, err := net.Dial("tcp", host)
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	conntools.Run(conn)
}
