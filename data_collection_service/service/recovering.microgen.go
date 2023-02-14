// Code generated by microgen 1.0.4. DO NOT EDIT.

package service

import (
	"context"
	service "esgrs/data_collection_service"
	types "esgrs/pkg/types"
	"fmt"
	log "github.com/go-kit/kit/log"
)

// RecoveringMiddleware recovers panics from method calls, writes to provided logger and returns the error of panic as method error.
func RecoveringMiddleware(logger log.Logger) Middleware {
	return func(next service.DataCollectionService) service.DataCollectionService {
		return &recoveringMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type recoveringMiddleware struct {
	logger log.Logger
	next   service.DataCollectionService
}

func (M recoveringMiddleware) GetCompanyByID(ctx context.Context, id int) (company types.Company, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "GetCompanyByID", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.GetCompanyByID(ctx, id)
}

func (M recoveringMiddleware) GetRawData(ctx context.Context, companyID int) (data []types.CriteriaRawData, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "GetRawData", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.GetRawData(ctx, companyID)
}

func (M recoveringMiddleware) GetCompanyList(ctx context.Context) (companies []types.Company, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "GetCompanyList", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.GetCompanyList(ctx)
}
