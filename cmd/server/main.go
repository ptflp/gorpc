package main

import (
	"github.com/ptflp/gorpc/repository"
	"github.com/ptflp/gorpc/rpca"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	pageRepository := repository.NewPageRepository()
	rpcAPI := rpca.NewAPI(pageRepository)
	err := rpc.Register(rpcAPI)

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
