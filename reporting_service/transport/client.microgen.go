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
	return EndpointsSet{GenerateReportEndpoint: opentracing.TraceClient(tracer, "GenerateReport")(endpoints.GenerateReportEndpoint)}
}

func (set EndpointsSet) GenerateReport(arg0 context.Context, arg1 types.ESGRatingResult) (res0 types.Report, res1 error) {
	request := GenerateReportRequest{Result: arg1}
	response, res1 := set.GenerateReportEndpoint(arg0, &request)
	if res1 != nil {
		return
	}
	return response.(*GenerateReportResponse).Rep, res1
}