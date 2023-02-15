package criteria_repository

import "esgrs/pkg/types"

type CriteriaRepository interface {
	GetCriteria(id int) (*types.CriteriaData, error)
	CreateCriteria(criteria *types.CriteriaData) (int, error)
	UpdateCriteria(criteria *types.CriteriaData) error
	DeleteCriteria(id int) error
}
