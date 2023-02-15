package category_repository

import (
	"esgrs/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CategoryRepositoryPostgres struct {
	db *gorm.DB
}

func (c CategoryRepositoryPostgres) GetAll() ([]*types.Category, error) {
	var categories []*types.Category
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c CategoryRepositoryPostgres) GetByID(id int) (*types.Category, error) {
	var category types.Category
	if err := c.db.Where("id =?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepositoryPostgres) GetByName(name string) (*types.Category, error) {
	var category types.Category
	if err := c.db.Where("name =?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepositoryPostgres) Create(category *types.Category) (*types.Category, error) {
	if err := c.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c CategoryRepositoryPostgres) Update(category *types.Category) (*types.Category, error) {
	if err := c.db.Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func NewCategoryRepositoryPostgres(connectURI string) *CategoryRepositoryPostgres {
	db, err := gorm.Open(postgres.Open(connectURI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&types.Category{}); err != nil {
		return nil
	}

	return &CategoryRepositoryPostgres{db: db}
}
