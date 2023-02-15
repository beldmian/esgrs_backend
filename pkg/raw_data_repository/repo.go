package raw_data_repository

import "esgrs/pkg/types"

type RawDataRepository interface {
	GetRawDataByCompanyID(companyID int) ([]*types.CriteriaRawData, error)
	CreateRawData(rawData *types.CriteriaRawData) (int, error)
}
