package main

import (
	"log"
	"net/http"

	"github.com/ProImpact/fakeapi/parser"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const prefix = "api"

func main() {
	r, err := parser.Open("./parser/example.json")
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// Middleware CORS simple para desarrollo
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	parser.AddRoutes(prefix, router, *r)
	log.Fatal(http.ListenAndServe(":4000", router))
}
