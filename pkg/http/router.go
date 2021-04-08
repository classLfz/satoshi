package http

import (
	"github.com/gorilla/mux"

	deviceCallerRouter "github.com/classlfz/satoshi/pkg/http/deviceCaller"
)

// NewAPIMux initial apis
func NewAPIMux() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/v1").Subrouter()
	initDeviceCallersAPI(s)

	return r
}

func initDeviceCallersAPI(r *mux.Router) {
	r.HandleFunc("/deviceCallers", deviceCallerRouter.CreateDeviceCaller).Methods("POST")
}
