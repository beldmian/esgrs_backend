package data_processing_service

import (
	"context"
	"esgrs/pkg/types"
)

// DataProcessingService defines service for processing data
// @microgen middleware, logging, tracing, metrics, recovering, error-logging, http
type DataProcessingService interface {
	// GetProcessedData accepts criteria data array & criteria metric array and returns the result
	GetProcessedData(ctx context.Context, companyID int) (result types.ESGRatingResult, err error)
}
