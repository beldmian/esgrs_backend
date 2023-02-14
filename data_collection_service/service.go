package data_collection_service

import (
	"context"
	"esgrs/pkg/types"
)

// DataCollectionService defines service for collecting data
// @microgen middleware, logging, tracing, metrics, recovering, error-logging, http
type DataCollectionService interface {
	// GetCompanyByID returns company by ID
	GetCompanyByID(ctx context.Context, id int) (company types.Company, err error)

	// GetRawData accepts parameter of company to collect data from
	GetRawData(ctx context.Context, companyID int) (data []types.CriteriaRawData, err error)

	// GetCompanyList returns a list of companies to be processed
	GetCompanyList(ctx context.Context) (companies []types.Company, err error)
}
