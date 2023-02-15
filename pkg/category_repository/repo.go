package category_repository

import "esgrs/pkg/types"

type CategoryRepository interface {
	GetAll() ([]*types.Category, error)
	GetByID(id int) (*types.Category, error)
	GetByName(name string) (*types.Category, error)
	Create(*types.Category) (*types.Category, error)
	Update(*types.Category) (*types.Category, error)
}
