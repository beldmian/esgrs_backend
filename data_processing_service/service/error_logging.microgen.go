// Code generated by microgen 1.0.4. DO NOT EDIT.

package service

import (
	"context"
	service "esgrs/data_processing_service"
	types "esgrs/pkg/types"
	log "github.com/go-kit/kit/log"
)

// ErrorLoggingMiddleware writes to logger any error, if it is not nil.
func ErrorLoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.DataProcessingService) service.DataProcessingService {
		return &errorLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type errorLoggingMiddleware struct {
	logger log.Logger
	next   service.DataProcessingService
}

func (M errorLoggingMiddleware) ProcessData(ctx context.Context, data []types.CriteriaData) (result types.ESGRatingResult, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "ProcessData", "message", err)
		}
	}()
	return M.next.ProcessData(ctx, data)
}