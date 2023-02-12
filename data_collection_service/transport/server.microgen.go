// Code generated by microgen 1.0.4. DO NOT EDIT.

package transport

import (
	"context"
	datacollectionservice "esgrs/data_collection_service"
	endpoint "github.com/go-kit/kit/endpoint"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func Endpoints(svc datacollectionservice.DataCollectionService) EndpointsSet {
	return EndpointsSet{
		CollectDataEndpoint:    CollectDataEndpoint(svc),
		GetCompanyListEndpoint: GetCompanyListEndpoint(svc),
	}
}

// TraceServerEndpoints is used for tracing endpoints on server side.
func TraceServerEndpoints(endpoints EndpointsSet, tracer opentracinggo.Tracer) EndpointsSet {
	return EndpointsSet{
		CollectDataEndpoint:    opentracing.TraceServer(tracer, "CollectData")(endpoints.CollectDataEndpoint),
		GetCompanyListEndpoint: opentracing.TraceServer(tracer, "GetCompanyList")(endpoints.GetCompanyListEndpoint),
	}
}

func CollectDataEndpoint(svc datacollectionservice.DataCollectionService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*CollectDataRequest)
		res0, res1 := svc.CollectData(arg0, req.Company)
		return &CollectDataResponse{Data: res0}, res1
	}
}

func GetCompanyListEndpoint(svc datacollectionservice.DataCollectionService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		res0, res1 := svc.GetCompanyList(arg0)
		return &GetCompanyListResponse{Companies: res0}, res1
	}
}