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
	return EndpointsSet{
		CollectDataEndpoint:    opentracing.TraceClient(tracer, "CollectData")(endpoints.CollectDataEndpoint),
		GetCompanyListEndpoint: opentracing.TraceClient(tracer, "GetCompanyList")(endpoints.GetCompanyListEndpoint),
	}
}

func (set EndpointsSet) CollectData(arg0 context.Context, arg1 types.Company) (res0 []types.CriteriaData, res1 error) {
	request := CollectDataRequest{Company: arg1}
	response, res1 := set.CollectDataEndpoint(arg0, &request)
	if res1 != nil {
		return
	}
	return response.(*CollectDataResponse).Data, res1
}

func (set EndpointsSet) GetCompanyList(arg0 context.Context) (res0 []types.Company, res1 error) {
	request := GetCompanyListRequest{}
	response, res1 := set.GetCompanyListEndpoint(arg0, &request)
	if res1 != nil {
		return
	}
	return response.(*GetCompanyListResponse).Companies, res1
}
