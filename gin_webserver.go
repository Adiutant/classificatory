package main

import (
	"fmt"
	"gin_webserver/http_server"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	server, err := http_server.NewPayloadServer(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	server.SetRoutes()
	err = server.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
