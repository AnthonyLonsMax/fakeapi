package main

import (
	"log"
	"net/http"

	"github.com/ProImpact/fakeapi/parser"
)

const prefix = "api"

func main() {
	r, err := parser.Open("./parser/example.json")
	if err != nil {
		log.Fatal(err)
	}
	router := http.NewServeMux()
	parser.AddRoutes(prefix, router, *r)
	log.Fatal(http.ListenAndServe(":4000", router))
}
