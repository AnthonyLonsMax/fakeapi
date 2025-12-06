package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ProImpact/fakeapi/util"
)

func Post(apiData []map[string]any, w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(400)
		return
	}
	apiData = append(apiData, body)
	util.SendJson(apiData[len(apiData)-1], w)
}
