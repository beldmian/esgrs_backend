package int_service

import (
	"context"
	transporthttp "esgrs/data_collection_service/transport/http"
	"esgrs/pkg/types"
	"net/url"
)

type DataProcessingServiceImplementation struct{}

func (svc DataProcessingServiceImplementation) GetProcessedData(ctx context.Context, companyID int) (result types.ESGRatingResult, err error) {
	apiURL, _ := url.Parse("http://data_collection_service:8080/")
	client := transporthttp.NewHTTPClient(apiURL)
	company, err := client.GetCompanyByID(ctx, companyID)
	if err != nil {
		return result, err
	}
	category, err := client.GetCategoryData(ctx, company.CategoryID)
	if err != nil {
		return result, err
	}
	data, err := client.GetRawData(ctx, companyID)
	if err != nil {
		return result, err
	}
	ECount := 0
	SCount := 0
	GCount := 0
	for _, criteria := range data {
		scoreOut := criteria.Score
		for _, weight := range category.CriteriaWeights {
			if criteria.CriteriaID == weight.CriteriaID {
				result.Rating += float64(scoreOut) * weight.Weight
				result.CriteriaScores = append(result.CriteriaScores, types.CriteriaScore{
					Score:          scoreOut,
					CriteriaWeight: weight,
				})
				switch criteria.CriteriaType {
				case types.ECriteria:
					result.ERating += float64(scoreOut) * weight.Weight
					ECount += 1
				case types.SCriteria:
					result.SRating += float64(scoreOut) * weight.Weight
					SCount += 1
				case types.GCriteria:
					result.GRating += float64(scoreOut) * weight.Weight
					GCount += 1
				}
			}
		}
	}
	result.Rating /= float64(len(data))
	result.ERating /= float64(ECount)
	result.SRating /= float64(SCount)
	result.GRating /= float64(GCount)
	return result, nil
}
