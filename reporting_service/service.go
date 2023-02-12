package reporting_service

import (
	"context"
	"esgrs/pkg/types"
)

// ReportingService defines service for generating reports
// @microgen middleware, logging, tracing, metrics, recovering, error-logging, http
type ReportingService interface {
	// GenerateReport accepts the result of ESGRatingResult and generates a report
	GenerateReport(ctx context.Context, result types.ESGRatingResult) (rep types.Report, err error)
}
