package data_collection_service

import (
	"context"
	"esgrs/pkg/types"
)

// DataCollectionService defines service for collecting data
// @microgen middleware, logging, tracing, metrics, recovering, error-logging, http
type DataCollectionService interface {
	// CollectData accepts parameter of company to collect data from
	CollectData(ctx context.Context, company types.Company) (data []types.CriteriaData, err error)

	// GetCompanyList returns a list of companies to be processed
	GetCompanyList(ctx context.Context) (companies []types.Company, err error)
}
