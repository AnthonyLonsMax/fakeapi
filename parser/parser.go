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
)

func Open(path string) (*model.Route, error) {
	var payload *model.Route
	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fileData, &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

var methods = []string{
	"GET", "POST", "PUT", "DELETE", "PATCH",
}

func AddRoutes(prefix string, router *http.ServeMux, route model.Route) {
	for routePath := range route {
		for _, method := range methods {
			path := fmt.Sprintf("%s /%s/%s", method, prefix, routePath)
			switch method {
			case "GET":
				router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
					handlers.Get(route[routePath], w, r)
				})
				log.Println(path)
				router.HandleFunc(path+"/{id}", func(w http.ResponseWriter, r *http.Request) {
					handlers.GetID(route[routePath], w, r)
				})
				log.Println(path + "/{id}")
			case "POST":
				router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
					handlers.Post(route[routePath], w, r)
				})
				log.Println(path)
			case "PATCH":
				router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
					handlers.Patch(route[routePath], w, r)
				})
				log.Println(path + "/{id}")
			case "DELETE":
				router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
					handlers.Delete(route[routePath], w, r)
				})
				log.Println(path + "/{id}")
			}
		}
	}
}
