// Code generated by microgen 1.0.4. DO NOT EDIT.

package transporthttp

import (
	transport "esgrs/reporting_service/transport"
	log "github.com/go-kit/kit/log"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	opentracinggo "github.com/opentracing/opentracing-go"
	http1 "net/http"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, logger log.Logger, tracer opentracinggo.Tracer, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/generate-report").Handler(
		http.NewServer(
			endpoints.GenerateReportEndpoint,
			_Decode_GenerateReport_Request,
			_Encode_GenerateReport_Response,
			append(opts, http.ServerBefore(
				opentracing.HTTPToContext(tracer, "GenerateReport", logger)))...))
	return mux
}
