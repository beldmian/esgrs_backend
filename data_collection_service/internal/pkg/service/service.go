package int_service

import (
	"context"
	"esgrs/pkg/types"
)

type DataCollectionServiceImplementation struct{}

func (svc DataCollectionServiceImplementation) CollectData(ctx context.Context, company types.Company) (data []types.CriteriaData, err error) {
	//TODO implement me
	panic("implement me")
}

func (svc DataCollectionServiceImplementation) GetCompanyList(ctx context.Context) (companies []types.Company, err error) {
	//TODO implement me
	panic("implement me")
}
