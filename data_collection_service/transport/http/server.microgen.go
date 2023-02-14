// Code generated by microgen 1.0.4. DO NOT EDIT.

package transporthttp

import (
	transport "esgrs/data_collection_service/transport"
	log "github.com/go-kit/kit/log"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	opentracinggo "github.com/opentracing/opentracing-go"
	http1 "net/http"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/get-company-byid").Handler(
		http.NewServer(
			endpoints.GetCompanyByIDEndpoint,
			_Decode_GetCompanyByID_Request,
			_Encode_GetCompanyByID_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "GetCompanyByID", logger)))...))
	mux.Methods("POST").Path("/get-raw-data").Handler(
		http.NewServer(
			endpoints.GetRawDataEndpoint,
			_Decode_GetRawData_Request,
			_Encode_GetRawData_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "GetRawData", logger)))...))
	mux.Methods("POST").Path("/get-company-list").Handler(
		http.NewServer(
			endpoints.GetCompanyListEndpoint,
			_Decode_GetCompanyList_Request,
			_Encode_GetCompanyList_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "GetCompanyList", logger)))...))
	return mux
}
