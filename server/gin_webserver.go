package main

import (
	"fmt"
	"gin_webserver/http_server"
)

func main() {
	server, err := http_server.NewPayloadServer("tcp", "pythonapp:777")
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
