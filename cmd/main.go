package main

import (
	"../repository"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	pageRepository := repository.NewPageRepository()
	fmt.Println(pageRepository)
	err := rpc.Register(pageRepository)

	if err != nil {
		log.Fatal("error registering rpc", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("serving rpc on TCP port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving: ", err)
	}

}
