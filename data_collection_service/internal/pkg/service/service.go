package int_service

import (
	"context"
	"esgrs/pkg/category_repository"
	"esgrs/pkg/company_repository"
	"esgrs/pkg/raw_data_repository"
	"esgrs/pkg/types"
)

type DataCollectionServiceImplementation struct {
	categoryRepository category_repository.CategoryRepository
	companyRepository  company_repository.CompanyRepository
	rawDataRepository  raw_data_repository.RawDataRepository
}

func NewDataCollectionService(postgresConnectionURI string) *DataCollectionServiceImplementation {
	return &DataCollectionServiceImplementation{
		categoryRepository: category_repository.NewCategoryRepositoryPostgres(postgresConnectionURI),
		companyRepository:  company_repository.NewCompanyRepositoryPostgres(postgresConnectionURI),
		rawDataRepository:  raw_data_repository.NewRawDataRepositoryPostgres(postgresConnectionURI),
	}
}

func (svc DataCollectionServiceImplementation) GetCategories(ctx context.Context) (categories []types.Category, err error) {
	categoriesPointers, err := svc.categoryRepository.GetAll()
	if err != nil {
		return nil, err
	}
	for _, category := range categoriesPointers {
		categories = append(categories, *category)
	}
	return
}

func (svc DataCollectionServiceImplementation) CreateCategory(ctx context.Context, category types.Category) (id int, err error) {
	outCategory, err := svc.categoryRepository.Create(&category)
	if err != nil {
		return -1, err
	}
	return outCategory.ID, nil
}

func (svc DataCollectionServiceImplementation) CreateCompany(ctx context.Context, company types.Company) (id int, err error) {
	outCompany, err := svc.companyRepository.Create(&company)
	if err != nil {
		return -1, err
	}
	return outCompany.ID, nil
}

func (svc DataCollectionServiceImplementation) GetCategoryData(ctx context.Context, categoryID int) (category types.Category, err error) {
	categoryDB, err := svc.categoryRepository.GetByID(categoryID)
	if err != nil {
		return types.Category{}, err
	}
	return *categoryDB, nil
}

func (svc DataCollectionServiceImplementation) GetCompanyByID(ctx context.Context, id int) (company types.Company, err error) {
	companyDB, err := svc.companyRepository.GetById(id)
	if err != nil {
		return types.Company{}, err
	}
	return *companyDB, nil
}

func (svc DataCollectionServiceImplementation) GetRawData(ctx context.Context, companyID int) (data []types.CriteriaRawData, err error) {
	dataPointers, err := svc.rawDataRepository.GetRawDataByCompanyID(companyID)
	if err != nil {
		return nil, err
	}
	for _, dataPointer := range dataPointers {
		data = append(data, *dataPointer)
	}
	return
}

func (svc DataCollectionServiceImplementation) GetCompanyList(ctx context.Context) (companies []types.Company, err error) {
	companiesPointers, err := svc.companyRepository.GetAll()
	if err != nil {
		return nil, err
	}
	for _, company := range companiesPointers {
		companies = append(companies, *company)
	}
	return
}
