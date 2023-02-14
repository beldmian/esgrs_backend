package visualization_service

import (
	"context"
	"esgrs/pkg/types"
)

// VisualizationService defines service for generating visualizaion
// @microgen middleware, logging, tracing, metrics, recovering, error-logging, http
type VisualizationService interface {
	// GenerateVisualization accepts the resultset of ESGRatingResult to generate a visualization
	GenerateVisualization(ctx context.Context, companyID int) (vis types.Visualization, err error)
}
