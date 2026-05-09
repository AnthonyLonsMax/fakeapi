package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/ProImpact/fakeapi/internal/parser"
	"github.com/ProImpact/fakeapi/internal/server"
	"github.com/go-chi/chi/v5"
)

var (
	port int
	file string
)

const prefix = "api"

func main() {
	flag.IntVar(&port, "port", 8080, "Server port for accept connections")
	flag.IntVar(&port, "p", 8080, "Server port for accept connections")
	flag.StringVar(&file, "file", "fakeapi.json", "Server data for generate the endpoints")
	flag.StringVar(&file, "f", "fakeapi.json", "Server data for generate the endpoints")
	flag.Parse()
	r, err := parser.Open(file)
	if err != nil {
		flag.PrintDefaults()
		log.Fatal(err)
	}
	router := chi.NewRouter()
	server.AddMiddlewares(router)
	server.AddLogger(os.Stdout)
	parser.AddRoutes(prefix, router, r)
	slog.Info("Server started", "port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
