package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type errorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (err errorResponse) Error() string {
	return strconv.Itoa(err.StatusCode) + ": " + err.Message
}

func (c *Controller) writeErrorResponse(err *errorResponse, w http.ResponseWriter) {
	http.Error(w, err.Message, err.StatusCode)
}

type successResponse struct {
	Message string `json:"message"`
}

func (c *Controller) writeResponse(suc *successResponse, w http.ResponseWriter) {
	if responseErr := json.NewEncoder(w).Encode(suc); responseErr != nil {
		w.WriteHeader(500)
		return
	}
}

func (c *Controller) writeCustomResponse(r interface{}, w http.ResponseWriter) {
	if responseErr := json.NewEncoder(w).Encode(r); responseErr != nil {
		w.WriteHeader(500)
		return
	}
}
