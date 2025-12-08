package handlers

import (
	"net/http"
	"slices"

	"github.com/ProImpact/fakeapi/types"
	"github.com/ProImpact/fakeapi/util"
)

func Delete(apiData *types.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	id, shouldReturn := getPathID(r, w, apiData.Data[resource])
	if shouldReturn {
		return
	}
	apiData.Data[resource] = slices.Delete(apiData.Data[resource], id, id+1)
	util.SendJson(apiData, w, 200)
}
