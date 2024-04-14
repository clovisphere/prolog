package main

import (
	"log"

	"github.com/clovisphere/prolog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Print("starting server...")
	log.Fatal(srv.ListenAndServe())
}
