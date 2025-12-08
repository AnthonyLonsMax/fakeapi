package handlers

import (
	"net/http"

	"github.com/ProImpact/fakeapi/types"
	"github.com/ProImpact/fakeapi/util"
)

func Post(apiData *types.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	err := util.ValidateRequest(r, body)
	if err != nil {
		types.ValidationError(w, err, r.URL.Path)
		return
	}
	apiData.Data[resource] = append(apiData.Data[resource], body)
	util.SendJson(apiData.Data[resource][len(apiData.Data[resource])-1], w, 200)
}
