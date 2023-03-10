// Code generated by microgen 1.0.4. DO NOT EDIT.

package transporthttp

import (
	transport "esgrs/data_collection_service/transport"
	log "github.com/go-kit/kit/log"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	httpkit "github.com/go-kit/kit/transport/http"
	opentracinggo "github.com/opentracing/opentracing-go"
	"net/url"
)

func NewHTTPClient(u *url.URL, opts ...httpkit.ClientOption) transport.EndpointsSet {
	return transport.EndpointsSet{
		CreateCategoryEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_CreateCategory_Request,
			_Decode_CreateCategory_Response,
			opts...,
		).Endpoint(),
		CreateCompanyEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_CreateCompany_Request,
			_Decode_CreateCompany_Response,
			opts...,
		).Endpoint(),
		GetCategoriesEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetCategories_Request,
			_Decode_GetCategories_Response,
			opts...,
		).Endpoint(),
		GetCategoryDataEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetCategoryData_Request,
			_Decode_GetCategoryData_Response,
			opts...,
		).Endpoint(),
		GetCompanyByIDEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetCompanyByID_Request,
			_Decode_GetCompanyByID_Response,
			opts...,
		).Endpoint(),
		GetCompanyListEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetCompanyList_Request,
			_Decode_GetCompanyList_Response,
			opts...,
		).Endpoint(),
		GetRawDataEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetRawData_Request,
			_Decode_GetRawData_Response,
			opts...,
		).Endpoint(),
	}
}

func TracingHTTPClientOptions(tracer opentracinggo.Tracer, logger log.Logger) func([]httpkit.ClientOption) []httpkit.ClientOption {
	return func(opts []httpkit.ClientOption) []httpkit.ClientOption {
		return append(opts, httpkit.ClientBefore(
			opentracing.ContextToHTTP(tracer, logger),
		))
	}
}
