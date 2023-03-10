// Code generated by microgen 1.0.4. DO NOT EDIT.

package transport

import (
	"context"
	visualizationservice "esgrs/visualization_service"
	endpoint "github.com/go-kit/kit/endpoint"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func Endpoints(svc visualizationservice.VisualizationService) EndpointsSet {
	return EndpointsSet{GenerateVisualizationEndpoint: GenerateVisualizationEndpoint(svc)}
}

// TraceServerEndpoints is used for tracing endpoints on server side.
func TraceServerEndpoints(endpoints EndpointsSet, tracer opentracinggo.Tracer) EndpointsSet {
	return EndpointsSet{GenerateVisualizationEndpoint: opentracing.TraceServer(tracer, "GenerateVisualization")(endpoints.GenerateVisualizationEndpoint)}
}

func GenerateVisualizationEndpoint(svc visualizationservice.VisualizationService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*GenerateVisualizationRequest)
		res0, res1 := svc.GenerateVisualization(arg0, req.CompanyID)
		return &GenerateVisualizationResponse{Vis: res0}, res1
	}
}
