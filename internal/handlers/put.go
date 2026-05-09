package handlers

import (
	"net/http"

	"github.com/ProImpact/fakeapi/internal/types"
	"github.com/ProImpact/fakeapi/pkg"
)

func Put(apiData *types.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	id, shouldReturn := getPathID(r, w, apiData.Data[resource])
	if shouldReturn {
		return
	}
	var body map[string]any
	err := pkg.ValidateRequest(r, body)
	if err != nil {
		types.ValidationError(w, err, r.URL.Path)
		return
	}
	apiData.Data[resource][id] = body
	pkg.SendJson(apiData.Data[resource][id], w, 200)
}
