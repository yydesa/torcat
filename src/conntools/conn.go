package conntools

import (
	"net"
	"os"
)

func output(conn net.Conn) {
	rb := make([]byte, 256)
	for {
		n, err := os.Stdin.Read(rb)
		if err != nil {
			panic("failed to read stdin: " + err.Error())
		}
		z := 0
		for z < n {
			m, err := conn.Write(rb[z:n])
			if err != nil {
				panic("failed to write: " + err.Error())
			}
			z += m
		}
	}
}

func Run(conn net.Conn) {
	go output(conn)

	rb := make([]byte, 256)
	for {
		n, err := conn.Read(rb)
		if err != nil {
			panic("failed to read: " + err.Error())
		}
		z := 0
		for z < n {
			m, err := os.Stdout.Write(rb[z:n])
			if err != nil {
				panic("failed to write stdout: " + err.Error())
			}
			z += m
		}
	}
	conn.Close()
}
