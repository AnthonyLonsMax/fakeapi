package handlers

import (
	"net/http"

	"github.com/ProImpact/fakeapi/model"
	"github.com/ProImpact/fakeapi/util"
)

func Patch(apiData *model.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	id, shouldReturn := getPathID(r, w, apiData.Data[resource])
	if shouldReturn {
		return
	}
	var body map[string]any
	err := util.ValidateRequest(r, body)
	if err != nil {
		model.ValidationError(w, err, r.URL.Path)
		return
	}
	util.PartialUpdate(body, apiData.Data[resource][id])
	util.SendJson(apiData.Data[resource][id], w, 200)
}
