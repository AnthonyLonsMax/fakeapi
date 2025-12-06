package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ProImpact/fakeapi/util"
)

func Patch(apiData []map[string]any, w http.ResponseWriter, r *http.Request) {
	id, shouldReturn := getPathID(r, w, apiData)
	if shouldReturn {
		return
	}
	var body map[string]any
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(400)
		return
	}
	util.PartialUpdate(body, apiData[id])
	util.SendJson(apiData[id], w)
}
