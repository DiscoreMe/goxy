package main

import (
	"github.com/DiscoreMe/goxy/server"
)

func main() {
	const hostMain = "127.0.0.1:5555"
	serv := server.Server{}
	if err := serv.Listen(hostMain); err != nil {
		panic(err)
	}
}
