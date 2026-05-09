package server

import (
	"net/http"
	"time"

	"github.com/ProImpact/fakeapi/internal/types"
	"github.com/ProImpact/fakeapi/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AddMiddlewares(router *chi.Mux) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
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
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		pkg.SendJson(&types.RequestErr{
			Code:      types.RESOURCE_NOT_FOUND,
			Message:   "Page not found",
			TimeStamp: time.Now(),
			Path:      r.URL.Path,
			Status:    http.StatusNotFound,
			Fault:     "client",
		}, w, 404)
	})
}
