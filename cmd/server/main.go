package main

import (
	"log"
	"net/http"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler"
	"github.com/tuanyang-qiu/ogen-demo/internal/oas"
)

func main() {

	srv, err := oas.NewServer(handler.NewProductHandler())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", srv))
}
