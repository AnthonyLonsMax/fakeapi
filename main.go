package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/ProImpact/fakeapi/parser"
	"github.com/ProImpact/fakeapi/server"
	"github.com/go-chi/chi/v5"
)

var (
	port = flag.Int("port", 8080, "Server port for accept connections")
	file = flag.String("file", "fakeapi.json", "Server data for generate the endpoints")
)

const prefix = "api"

func main() {
	flag.Parse()
	r, err := parser.Open(*file)
	if err != nil {
		flag.PrintDefaults()
		log.Fatal(err)
	}
	router := chi.NewRouter()
	server.AddMiddlewares(router)
	server.AddLogger(os.Stdout)
	parser.AddRoutes(prefix, router, r)
	slog.Info("Server started", "port", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
