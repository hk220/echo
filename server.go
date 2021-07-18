package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func echo(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadBytes(byte('\n'))
		switch err {
		case nil:
			break
		case io.EOF:
		default:
			fmt.Println(fmt.Errorf("ERROR: %w\n", err))
		}
		conn.Write(line)
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:9001")
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR: %w\n", err))
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(fmt.Errorf("ERROR: %w\n", err))
			continue
		}
		go echo(conn)
	}
}
