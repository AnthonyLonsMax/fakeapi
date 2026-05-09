package handlers

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/ProImpact/fakeapi/internal/types"
	"github.com/ProImpact/fakeapi/pkg"
)

func invalidValueResponse(w http.ResponseWriter, r *http.Request, err error) {
	pkg.SendJson(&types.RequestErr{
		Code:      types.INVALID_VALUE,
		Message:   err.Error(),
		TimeStamp: time.Now(),
		Path:      r.URL.Path,
		Status:    http.StatusBadRequest,
		Fault:     "client",
	}, w, 400)
}

func Get(apiData *types.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	limit, err := pkg.GetIntFromQuery(r, "limit", 10)
	if err != nil {
		invalidValueResponse(w, r, err)
		return
	}
	log.Println("Query param", "limit", limit)
	offset, err := pkg.GetIntFromQuery(r, "offset", 0)
	if err != nil {
		invalidValueResponse(w, r, err)
		return
	}
	log.Println("Query param", "offset", offset)
	if offset < 0 || offset > len(apiData.Data[resource])-1 {
		pkg.SendJson(&types.RequestErr{
			Code:      types.OUT_OF_RANGE,
			Message:   fmt.Sprintf("Offset should be greater than 0 and lower than %d", len(apiData.Data)),
			TimeStamp: time.Now(),
			Path:      r.URL.Path,
			Status:    http.StatusBadRequest,
			Fault:     "client",
		}, w, 400)
		return
	}
	sortKey := pkg.GetStringFromQuery(r, "sort", "key")
	arrayCopy := slices.Clone(apiData.Data[resource])
	pkg.SortMap(arrayCopy, sortKey)
	if offset+limit >= len(apiData.Data[resource])-1 {
		pkg.SendJson(arrayCopy, w, 200)
		return
	}
	pkg.SendJson(arrayCopy[offset:limit], w, 200)
}

func GetID(apiData *types.ApiData, resource string, w http.ResponseWriter, r *http.Request) {
	id, shouldReturn := getPathID(r, w, apiData.Data[resource])
	if shouldReturn {
		return
	}
	pkg.SendJson(apiData.Data[resource][id], w, 200)
}

func getPathID(r *http.Request, w http.ResponseWriter, apiData []map[string]any) (int, bool) {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		pkg.SendJson(&types.RequestErr{
			Code:      types.INVALID_ARGUMENT,
			Message:   err.Error(),
			TimeStamp: time.Now(),
			Path:      r.URL.Path,
			Status:    http.StatusBadRequest,
			Fault:     "client",
		}, w, 400)
		return 0, true
	}
	if id > len(apiData) || id == 0 {
		pkg.SendJson(&types.RequestErr{
			Code:      types.INVALID_ARGUMENT,
			Message:   "index out of range",
			TimeStamp: time.Now(),
			Path:      r.URL.Path,
			Status:    http.StatusBadRequest,
			Fault:     "client",
		}, w, 404)
		return 0, true
	}
	return id - 1, false
}
