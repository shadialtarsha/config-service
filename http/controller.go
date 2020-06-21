package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shadialtarsha/config-service/service"
)

// Controller exports the handlers for the endpoints
type Controller struct {
	Config *service.Config
}

// ReadConfig writes the config for the given service to the ResponseWriter.
func (c *Controller) ReadConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)

	serviceName, ok := vars["serviceName"]
	if !ok {
		c.writeErrorResponse(&errorResponse{StatusCode: 400, Message: "Service name is required"}, w)
		return
	}

	env := r.URL.Query().Get("env")
	if env == "" {
		env = "development"
	}

	config, err := c.Config.Get(env, serviceName)
	if err != nil {
		c.writeErrorResponse(&errorResponse{StatusCode: 500, Message: "Something went wrong"}, w)
		return
	}

	c.writeCustomResponse(config, w)
}
