package handlers

import (
	"net/http"
	"strconv"

	"github.com/ProImpact/fakeapi/model"
	"github.com/ProImpact/fakeapi/util"
)

func Get(apiData *model.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	util.SendJson(apiData.Data[resource], w)
}

func GetID(apiData *model.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	id, shouldReturn := getPathID(r, w, apiData.Data[resource])
	if shouldReturn {
		return
	}
	util.SendJson(apiData.Data[resource][id-1], w)
}

func getPathID(r *http.Request, w http.ResponseWriter, apiData []map[string]any) (int, bool) {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return 0, true
	}
	if id > len(apiData) || id == 0 {
		w.WriteHeader(404)
		w.Write([]byte("Index out of range"))
		return 0, true
	}
	return id - 1, false
}
