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

	// GetCategories returns a list of categories to be processed
	GetCategories(ctx context.Context) (categories []types.Category, err error)

	// GetCategoryData returns data of category
	GetCategoryData(ctx context.Context, categoryID int) (category types.Category, err error)

	// CreateCategory creates a new category
	CreateCategory(ctx context.Context, category types.Category) (id int, err error)

	// CreateCompany creates a new company
	CreateCompany(ctx context.Context, company types.Company) (id int, err error)
}
