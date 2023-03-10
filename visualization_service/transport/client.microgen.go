// Code generated by microgen 1.0.4. DO NOT EDIT.

package transport

import (
	"context"
	types "esgrs/pkg/types"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	opentracinggo "github.com/opentracing/opentracing-go"
)

// TraceClientEndpoints is used for tracing endpoints on client side.
func TraceClientEndpoints(endpoints EndpointsSet, tracer opentracinggo.Tracer) EndpointsSet {
	return EndpointsSet{GenerateVisualizationEndpoint: opentracing.TraceClient(tracer, "GenerateVisualization")(endpoints.GenerateVisualizationEndpoint)}
}

func (set EndpointsSet) GenerateVisualization(arg0 context.Context, arg1 int) (res0 types.Visualization, res1 error) {
	request := GenerateVisualizationRequest{CompanyID: arg1}
	response, res1 := set.GenerateVisualizationEndpoint(arg0, &request)
	if res1 != nil {
		return
	}
	return response.(*GenerateVisualizationResponse).Vis, res1
}
