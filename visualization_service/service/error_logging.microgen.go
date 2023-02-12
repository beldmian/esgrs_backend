// Code generated by microgen 1.0.4. DO NOT EDIT.

package service

import (
	"context"
	types "esgrs/pkg/types"
	service "esgrs/visualization_service"
	log "github.com/go-kit/kit/log"
)

// ErrorLoggingMiddleware writes to logger any error, if it is not nil.
func ErrorLoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.VisualizationService) service.VisualizationService {
		return &errorLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type errorLoggingMiddleware struct {
	logger log.Logger
	next   service.VisualizationService
}

func (M errorLoggingMiddleware) GenerateVisualization(ctx context.Context, result types.ESGRatingResult) (vis types.Visualization, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GenerateVisualization", "message", err)
		}
	}()
	return M.next.GenerateVisualization(ctx, result)
}
