package handlers

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/ProImpact/fakeapi/model"
	"github.com/ProImpact/fakeapi/util"
)

func invalidValueResponse(w http.ResponseWriter, r *http.Request, err error) {
	util.SendJson(&model.RequestErr{
		Code:      model.INVALID_VALUE,
		Message:   err.Error(),
		TimeStamp: time.Now(),
		Path:      r.URL.Path,
		Status:    http.StatusBadRequest,
		Fault:     "client",
	}, w, 400)
}

func Get(apiData *model.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	limit, err := util.GetIntFromQuery(r, "limit", 10)
	if err != nil {
		invalidValueResponse(w, r, err)
		return
	}
	log.Println("Query param", "limit", limit)
	offset, err := util.GetIntFromQuery(r, "offset", 0)
	if err != nil {
		invalidValueResponse(w, r, err)
		return
	}
	log.Println("Query param", "offset", offset)
	if offset < 0 || offset > len(apiData.Data[resource])-1 {
		util.SendJson(&model.RequestErr{
			Code:      model.OUT_OF_RANGE,
			Message:   fmt.Sprintf("Offset should be greater than 0 and lower than %d", len(apiData.Data)),
			TimeStamp: time.Now(),
			Path:      r.URL.Path,
			Status:    http.StatusBadRequest,
			Fault:     "client",
		}, w, 400)
		return
	}
	sortKey := util.GetStringFromQuery(r, "sort", "key")
	arrayCopy := slices.Clone(apiData.Data[resource])
	util.SortMap(arrayCopy, sortKey)
	if offset+limit >= len(apiData.Data[resource])-1 {
		util.SendJson(arrayCopy, w, 200)
		return
	}
	util.SendJson(arrayCopy[offset:limit], w, 200)
}

func GetID(apiData *model.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	id, shouldReturn := getPathID(r, w, apiData.Data[resource])
	if shouldReturn {
		return
	}
	util.SendJson(apiData.Data[resource][id], w, 200)
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
