package int_service

import (
	"context"
	"esgrs/pkg/types"
)

type DataProcessingServiceImplementation struct{}

func (svc DataProcessingServiceImplementation) ProcessData(ctx context.Context, data []types.CriteriaData) (result types.ESGRatingResult, err error) {
	//TODO implement me
	panic("implement me")
}
