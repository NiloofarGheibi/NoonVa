package main

import (
	"github.com/gotschmarcel/goserv"
	"log"
	"net/http"
)

func main() {

	server := goserv.NewServer()

	// Handle Get Request
	server.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if err := goserv.WriteJSON(w, "Welcome to NoonVa!"); err != nil {
			goserv.Context(r).Error(err, http.StatusInternalServerError)
			return
		}
	})

	log.Fatalln(server.Listen(":60234"))
}
