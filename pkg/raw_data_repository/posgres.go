package raw_data_repository

import (
	"esgrs/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RawDataRepositoryPostgres struct {
	db *gorm.DB
}

func (r RawDataRepositoryPostgres) CreateRawData(rawData *types.CriteriaRawData) (int, error) {
	err := r.db.Create(rawData).Error
	return rawData.ID, err
}

func (r RawDataRepositoryPostgres) GetRawDataByCompanyID(companyID int) ([]*types.CriteriaRawData, error) {
	var rawDatas []*types.CriteriaRawData
	err := r.db.Where("company_id =?", companyID).Find(&rawDatas).Error
	return rawDatas, err
}

func NewRawDataRepositoryPostgres(connectionURI string) *RawDataRepositoryPostgres {
	db, err := gorm.Open(postgres.Open(connectionURI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&types.CriteriaRawData{}); err != nil {
		return nil
	}

	return &RawDataRepositoryPostgres{db: db}
}
