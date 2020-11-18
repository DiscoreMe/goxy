package server

import (
	"bytes"
	"io"
	"net"
	"time"
)

type Tunnel struct {
	host     string
	owner    *net.Conn
	listener net.Listener
}

func newTunnel(network, host string) (Tunnel, error) {
	listener, err := net.Listen(network, host)
	if err != nil {
		return Tunnel{}, err
	}

	tunnel := Tunnel{
		host:     host,
		listener: listener,
	}

	go func() {
		tunnel.listen()
	}()

	return tunnel, nil
}

func (t *Tunnel) listen() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			return
		}
		if t.owner == nil {
			t.owner = &conn

			for {
				(*t.owner).Write([]byte("Hello"))
				time.Sleep(1 * time.Second)
			}
		} else {
			var buf bytes.Buffer
			io.Copy(&buf, conn)
			io.Copy(*t.owner, &buf)
		}
	}
}
