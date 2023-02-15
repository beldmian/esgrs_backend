package company_repository

import "esgrs/pkg/types"

type CompanyRepository interface {
	GetAll() ([]*types.Company, error)
	GetByName(name string) (*types.Company, error)
	GetById(id int) (*types.Company, error)
	GetByIds(ids ...int) ([]*types.Company, error)
	Create(company *types.Company) (*types.Company, error)
	Update(company *types.Company) (*types.Company, error)
	Delete(id int) error
}
