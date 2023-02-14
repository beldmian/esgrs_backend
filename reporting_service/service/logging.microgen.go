// Code generated by microgen 1.0.4. DO NOT EDIT.

package service

import (
	"context"
	types "esgrs/pkg/types"
	service "esgrs/reporting_service"
	log "github.com/go-kit/kit/log"
	"time"
)

// LoggingMiddleware writes params, results and working time of method call to provided logger after its execution.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.ReportingService) service.ReportingService {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   service.ReportingService
}

func (M loggingMiddleware) GenerateReport(arg0 context.Context, arg1 int) (res0 types.Report, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GenerateReport",
			"message", "GenerateReport called",
			"request", logGenerateReportRequest{CompanyID: arg1},
			"response", logGenerateReportResponse{Rep: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GenerateReport(arg0, arg1)
}

type (
	logGenerateReportRequest struct {
		CompanyID int
	}
	logGenerateReportResponse struct {
		Rep types.Report
	}
)
