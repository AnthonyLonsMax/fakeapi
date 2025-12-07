package handlers

import (
	"net/http"

	"github.com/ProImpact/fakeapi/model"
	"github.com/ProImpact/fakeapi/util"
)

func Post(apiData *model.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	err := util.ValidateRequest(r, body)
	if err != nil {
		model.ValidationError(w, err, r.URL.Path)
		return
	}
	apiData.Data[resource] = append(apiData.Data[resource], body)
	util.SendJson(apiData.Data[resource][len(apiData.Data[resource])-1], w, 200)
}
