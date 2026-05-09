package handlers

import (
	"net/http"

	"github.com/ProImpact/fakeapi/internal/types"
	"github.com/ProImpact/fakeapi/pkg"
)

func Post(apiData *types.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	err := pkg.ValidateRequest(r, body)
	if err != nil {
		types.ValidationError(w, err, r.URL.Path)
		return
	}
	apiData.Data[resource] = append(apiData.Data[resource], body)
	pkg.SendJson(apiData.Data[resource][len(apiData.Data[resource])-1], w, 200)
}
