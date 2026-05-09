package handlers

import (
	"net/http"
	"slices"

	"github.com/ProImpact/fakeapi/internal/types"
	"github.com/ProImpact/fakeapi/pkg"
)

func Delete(apiData *types.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	id, shouldReturn := getPathID(r, w, apiData.Data[resource])
	if shouldReturn {
		return
	}
	apiData.Data[resource] = slices.Delete(apiData.Data[resource], id, id+1)
	pkg.SendJson(apiData, w, 200)
}
