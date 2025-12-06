package handlers

import (
	"net/http"
	"slices"

	"github.com/ProImpact/fakeapi/util"
)

func Delete(apiData []map[string]any, w http.ResponseWriter, r *http.Request) {
	id, shouldReturn := getPathID(r, w, apiData)
	if shouldReturn {
		return
	}
	apiData = slices.Delete(apiData, id, id)
	util.SendJson(apiData, w)
}
