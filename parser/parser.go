// parser provides the json unmarshalling and code generation
package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ProImpact/fakeapi/handlers"
	"github.com/ProImpact/fakeapi/model"
	"github.com/go-chi/chi/v5"
)

func Open(path string) (*model.ApiData, error) {
	var payload model.Route
	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fileData, &payload)
	if err != nil {
		return nil, err
	}
	return &model.ApiData{
		Data: payload,
	}, nil
}

var methods = []string{
	"GET", "POST", "PUT", "DELETE", "PATCH", "PUT",
}

func AddRoutes(prefix string, router *chi.Mux, route *model.ApiData) {
	for routePath := range route.Data {
		for _, method := range methods {
			path := fmt.Sprintf("%s /%s/%s", method, prefix, routePath)
			switch method {
			case "GET":
				router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
					handlers.Get(route, routePath, w, r)
				})
				router.HandleFunc(path+"/{id}", func(w http.ResponseWriter, r *http.Request) {
					handlers.GetID(route, routePath, w, r)
				})
				log.Println(path)
				log.Println(path + "/{id}")
			case "POST":
				router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
					handlers.Post(route, routePath, w, r)
				})
				log.Println(path)
			case "PATCH":
				router.HandleFunc(path+"/{id}", func(w http.ResponseWriter, r *http.Request) {
					handlers.Patch(route, routePath, w, r)
				})
				log.Println(path + "/{id}")
			case "DELETE":
				router.HandleFunc(path+"/{id}", func(w http.ResponseWriter, r *http.Request) {
					handlers.Delete(route, routePath, w, r)
				})
				log.Println(path + "/{id}")
			case "PUT":
				router.HandleFunc(path+"/{id}", func(w http.ResponseWriter, r *http.Request) {
					handlers.Put(route, routePath, w, r)
				})
				log.Println(path + "/{id}")
			}
		}
	}
}
