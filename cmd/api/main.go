package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8060

func main() {
	// creating an oject of Type GQServer
	var serve GQServer

	// Starting Server
	log.Println("Starting Server on port: ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), serve.routes())
	if err != nil {
		log.Fatal(err)

	}
}
