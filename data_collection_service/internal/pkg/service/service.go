package int_service

import (
	"context"
	"esgrs/pkg/types"
	"time"
)

type DataCollectionServiceImplementation struct{}

func (svc DataCollectionServiceImplementation) GetCompanyByID(ctx context.Context, id int) (company types.Company, err error) {
	//TODO implement me
	panic("implement me")
}

func (svc DataCollectionServiceImplementation) GetRawData(ctx context.Context, companyID int) (data []types.CriteriaRawData, err error) {
	// TODO normal implementation
	out := []types.CriteriaRawData{
		{
			CompanyID:      companyID,
			CriteriaID:     0,
			Score:          3,
			CollectionDate: time.Now(),
		},
		{
			CompanyID:      companyID,
			CriteriaID:     1,
			Score:          2,
			CollectionDate: time.Now(),
		},
	}
	return out, nil
}

func (svc DataCollectionServiceImplementation) GetCompanyList(ctx context.Context) (companies []types.Company, err error) {
	//TODO implement me
	panic("implement me")
}
