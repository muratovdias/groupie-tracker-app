package main

import (
	"fmt"
	"log"
	"net/http"

	"tracker/internal/server"
)

const port = ":8081"

func main() {
	server := server.New()
	fmt.Printf("Starting server at port %s\nhttp://localhost%s/\n", port, port)
	if err := http.ListenAndServe(port, server.Route()); err != nil {
		log.Fatal()
	}
}
