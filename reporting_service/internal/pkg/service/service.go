package int_service

import (
	"context"
	"esgrs/pkg/types"
)

type ReportingServiceImplementation struct{}

func (svc ReportingServiceImplementation) GenerateReport(ctx context.Context, result types.ESGRatingResult) (rep types.Report, err error) {
	//TODO implement me
	panic("implement me")
}
