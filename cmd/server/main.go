package main

import (
	"log"
	"net/http"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler"
	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
)

func main() {

	srv, err := ogen.NewServer(handler.NewProductHandler())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", srv))
}
