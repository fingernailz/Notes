package main

import (
	"fmt"
	"notes/server"
	storage "notes/storage"
)

var SetOfBooks storage.Library = storage.Library{}

func main() {
	fmt.Println("Server starting")

	err := server.StartAPIService()
	if err != nil {
		fmt.Println(err)
	}
}
