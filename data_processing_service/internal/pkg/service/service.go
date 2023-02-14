package int_service

import (
	"context"
	transporthttp "esgrs/data_collection_service/transport/http"
	"esgrs/pkg/types"
	"log"
	"net/url"
)

type DataProcessingServiceImplementation struct{}

func (svc DataProcessingServiceImplementation) GetProcessedData(ctx context.Context, companyID int) (result types.ESGRatingResult, err error) {
	apiURL, _ := url.Parse("http://data_collection_service:8080/")
	client := transporthttp.NewHTTPClient(apiURL)
	data, err := client.GetRawData(ctx, companyID)
	log.Println(data)
	for _, criteria := range data {
		result.Rating += float64(criteria.Score)
	}
	result.Rating /= float64(len(data))
	return result, nil
}
