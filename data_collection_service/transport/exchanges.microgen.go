// Code generated by microgen 1.0.4. DO NOT EDIT.

package transport

import types "esgrs/pkg/types"

type (
	GetCompanyByIDRequest struct {
		Id int `json:"id"`
	}
	GetCompanyByIDResponse struct {
		Company types.Company `json:"company"`
	}

	GetRawDataRequest struct {
		CompanyID int `json:"company_id"`
	}
	GetRawDataResponse struct {
		Data []types.CriteriaRawData `json:"data"`
	}

	// Formal exchange type, please do not delete.
	GetCompanyListRequest  struct{}
	GetCompanyListResponse struct {
		Companies []types.Company `json:"companies"`
	}

	// Formal exchange type, please do not delete.
	GetCategoriesRequest  struct{}
	GetCategoriesResponse struct {
		Categories []types.Category `json:"categories"`
	}

	GetCategoryDataRequest struct {
		CategoryID int `json:"category_id"`
	}
	GetCategoryDataResponse struct {
		Category types.Category `json:"category"`
	}

	CreateCategoryRequest struct {
		Category types.Category `json:"category"`
	}
	CreateCategoryResponse struct {
		Id int `json:"id"`
	}

	CreateCompanyRequest struct {
		Company types.Company `json:"company"`
	}
	CreateCompanyResponse struct {
		Id int `json:"id"`
	}
)
