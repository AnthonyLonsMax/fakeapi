package types

import (
	"encoding/json"
	"net/http"
	"time"
)

type RequestErr struct {
	Code      ErrorCode      `json:"code,omitempty"`
	Message   string         `json:"message,omitempty"`
	Details   map[string]any `json:"details,omitempty"`
	TimeStamp time.Time      `json:"time_stamp,omitempty"`
	Path      string         `json:"path,omitempty"`
	Status    int            `json:"status,omitempty"`
	Fault     string         `json:"fault,omitempty"` // server | client
}

type SuccessResp struct {
	Data      any            `json:"data,omitempty"`
	Message   string         `json:"message,omitempty"`
	Meta      map[string]any `json:"meta,omitempty"`
	TimeStamp time.Time      `json:"time_stamp,omitempty"`
}

func ValidationError(w http.ResponseWriter, err error, path string) {
	json.NewEncoder(w).Encode(&RequestErr{
		Code:      VALIDATION_ERROR,
		Message:   err.Error(),
		TimeStamp: time.Now(),
		Path:      path,
		Status:    http.StatusBadRequest,
		Fault:     "client",
	})
}
