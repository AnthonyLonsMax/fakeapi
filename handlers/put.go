package handlers

import (
	"net/http"

	"github.com/ProImpact/fakeapi/types"
	"github.com/ProImpact/fakeapi/util"
)

func Put(apiData *types.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	id, shouldReturn := getPathID(r, w, apiData.Data[resource])
	if shouldReturn {
		return
	}
	var body map[string]any
	err := util.ValidateRequest(r, body)
	if err != nil {
		types.ValidationError(w, err, r.URL.Path)
		return
	}
	apiData.Data[resource][id] = body
	util.SendJson(apiData.Data[resource][id], w, 200)
}
