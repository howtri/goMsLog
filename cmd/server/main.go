package main

import (
	"log"

	"github.com/howtri/goMsLog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
