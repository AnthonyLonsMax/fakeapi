package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ProImpact/fakeapi/model"
	"github.com/ProImpact/fakeapi/util"
)

func Post(apiData *model.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(400)
		return
	}
	apiData.Data[resource] = append(apiData.Data[resource], body)
	util.SendJson(apiData.Data[resource][len(apiData.Data[resource])-1], w)
}
