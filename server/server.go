package server

import (
	"log"
	"net"
)

// todo: make random port
const host = "127.0.0.1:25078"
const networkDefault = "tcp"

type Server struct {
	tunnels []Tunnel
}

func (s *Server) Listen(host string) error {
	listener, err := net.Listen(networkDefault, host)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		var id = make([]byte, 1)
		conn.Read(id)

		if id[0] == 1 {
			host, err := s.createTunnel(networkDefault)
			if err != nil {
				log.Println(err)
				continue
			}
			conn.Write([]byte(host))
		}

		conn.Close()
	}
}

func (s *Server) createTunnel(network string) (string, error) {
	tunnel, err := newTunnel(network, host)
	if err != nil {
		return "", err
	}
	s.tunnels = append(s.tunnels, tunnel)

	return tunnel.host, nil
}
