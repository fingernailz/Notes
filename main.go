package main

import (
	"fmt"
	"notes/server"
)

func main() {
	fmt.Println("Server starting")

	err := server.StartAPIService()
	if err != nil {
		fmt.Println(err)
	}
}
