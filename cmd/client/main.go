package main

import (
	"bytes"
	"fmt"
	"io"
	"net"

	"github.com/DiscoreMe/goxy/protocol"
)

const urlMain = "127.0.0.1:5555"

func main() {
	conn, err := net.Dial("tcp", urlMain)
	if err != nil {
		panic(err)
	}

	if _, err := conn.Write(protocol.AuthProtocol()); err != nil {
		panic(err)
	}

	var host bytes.Buffer
	if _, err := io.Copy(&host, conn); err != nil {
		panic(err)
	}

	conn.Close()

	fmt.Println(host.String())

	for {
		conn, err = net.Dial("tcp", host.String())
		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		io.Copy(&buf, conn)
		if buf.Len() > 0 {
			fmt.Println(buf.String())
		}

		conn.Close()
	}

}
