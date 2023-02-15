package company_repository

import (
	"esgrs/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CompanyRepositoryPostgres struct {
	db *gorm.DB
}

func (c CompanyRepositoryPostgres) GetAll() ([]*types.Company, error) {
	var companies []*types.Company
	err := c.db.Find(&companies).Error
	return companies, err
}

func (c CompanyRepositoryPostgres) GetByName(name string) (*types.Company, error) {
	var companies []*types.Company
	err := c.db.Where("name =?", name).Find(&companies).Error
	return companies[0], err
}

func (c CompanyRepositoryPostgres) GetById(id int) (*types.Company, error) {
	var companies []*types.Company
	err := c.db.Where("id =?", id).Find(&companies).Error
	return companies[0], err

}

func (c CompanyRepositoryPostgres) GetByIds(ids ...int) ([]*types.Company, error) {
	var companies []*types.Company
	err := c.db.Where("id in (?)", ids).Find(&companies).Error
	return companies, err
}

func (c CompanyRepositoryPostgres) Create(company *types.Company) (*types.Company, error) {
	err := c.db.Create(company).Error
	return company, err
}

func (c CompanyRepositoryPostgres) Update(company *types.Company) (*types.Company, error) {
	err := c.db.Save(company).Error
	return company, err
}

func (c CompanyRepositoryPostgres) Delete(id int) error {
	err := c.db.Delete(&types.Company{}, id).Error
	return err
}

func NewCompanyRepositoryPostgres(connectionURI string) *CompanyRepositoryPostgres {
	db, err := gorm.Open(postgres.Open(connectionURI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&types.Company{}); err != nil {
		return nil
	}

	return &CompanyRepositoryPostgres{db: db}
}
