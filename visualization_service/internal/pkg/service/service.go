package int_service

import (
	"context"
	"esgrs/pkg/types"
)

type VisualizationServiceImplementation struct{}

func (svc VisualizationServiceImplementation) GenerateVisualization(ctx context.Context, companyID int) (vis types.Visualization, err error) {
	//TODO implement me
	panic("implement me")
}
