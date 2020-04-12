package main

import (
	"fmt"
	"github.com/ptflp/gorpc"
	"log"
	"net/rpc"
	"time"
)

func main() {
	var reply gorpc.Page
	var db []gorpc.Page

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := [3]gorpc.Page{
		{Title: "First", Body: "A first item"},
		{Title: "Second", Body: "A second item"},
		{Title: "Third", Body: "A third item"},
	}

	for i := range a {
		err = client.Call("API.AddItem", a[i], &reply)
		if err != nil {
			log.Println("something wrong on AddItem")
		}
	}

	client.Call("API.GetAll", "", &db)

	fmt.Println("Database: ", db)

	client.Call("API.EditItem", gorpc.Page{ID: 2, Title: "Second", Body: "A new second item"}, &reply)

	client.Call("API.DeleteItem", a[2], &reply)
	client.Call("API.GetAll", "", &db)
	fmt.Println("Database: ", db)

	client.Call("API.GetByName", "First", &reply)
	fmt.Println("first item: ", reply)
	time.Sleep(50 * time.Second)
	client.Call("API.GetByName", "First", &reply)
}
