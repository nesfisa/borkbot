package borkbot

import (
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer serves as the main route constructor for the bork API
func NewHTTPServer(endpoints Endpoints, logger kitlog.Logger) http.Handler {
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}
	r := mux.NewRouter()
	r.Handle("/bork", httptransport.NewServer(
		endpoints.FetchBorkEndpoint,
		decodeFetchBorkRequest,
		encodeResponse,
		opts...,
	)).Methods("POST")
	return r
}
