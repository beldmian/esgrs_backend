package criteria_repository

import (
	"esgrs/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CriteriaRepositoryPostgres struct {
	db *gorm.DB
}

func (c CriteriaRepositoryPostgres) GetCriteria(id int) (*types.CriteriaData, error) {
	var criteria types.CriteriaData
	if err := c.db.Where("id =?", id).Find(&criteria).Error; err != nil {
		return nil, err
	}
	return &criteria, nil
}

func (c CriteriaRepositoryPostgres) CreateCriteria(criteria *types.CriteriaData) (int, error) {
	if err := c.db.Create(criteria).Error; err != nil {
		return 0, err
	}
	return criteria.ID, nil
}

func (c CriteriaRepositoryPostgres) UpdateCriteria(criteria *types.CriteriaData) error {
	if err := c.db.Save(criteria).Error; err != nil {
		return err
	}
	return nil
}

func (c CriteriaRepositoryPostgres) DeleteCriteria(id int) error {
	if err := c.db.Delete(&types.CriteriaData{}, "id =?", id).Error; err != nil {
		return err
	}
	return nil
}

func NewCriteriaRepositoryPostgres(connectionURI string) *CriteriaRepositoryPostgres {
	db, err := gorm.Open(postgres.Open(connectionURI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&types.CriteriaData{}); err != nil {
		return nil
	}
	return &CriteriaRepositoryPostgres{db: db}
}
