package types

import (
	"gorm.io/gorm"
	"time"
)

type CriteriaType int

const (
	ECriteria CriteriaType = iota
	SCriteria
	GCriteria
)

type CriteriaID int

const (
	CriteriaGasEmissions CriteriaID = iota
	CriteriaEnergyConsumption
	CriteriaAirPollution
	CriteriaWaterManagement
	CriteriaWasteManagement
	CriteriaReclamation
	CriteriaSupplyChain
	CriteriaAccidents
	CriteriaEnvironmentalInvestments
	CriteriaPortfolioValuation

	CriteriaStaffMotivation
	CriteriaEqualityPolicity
	CriteriaHealthSafety
	CriteriaClientInteraction

	CriteriaOwnershipStructure
	CriteriaControls
	CriteriaStrategy
	CriteriaRewardSystem
	CriteriaRiskManagement
	CriteriaDisclosure
)

type CriteriaRawData struct {
	gorm.Model
	ID             int          `json:"id"`
	CriteriaID     CriteriaID   `json:"criteria_id"`
	CompanyID      int          `json:"company_id"`
	CriteriaType   CriteriaType `json:"criteria_type"`
	Score          int          `json:"score"`
	CollectionDate time.Time    `json:"collection_date"`
}

type CriteriaData struct {
	gorm.Model
	ID           int          `json:"id"`
	CriteriaID   CriteriaID   `json:"criteria_id,omitempty"`
	Name         string       `json:"name,omitempty"`
	Description  string       `json:"description,omitempty"`
	CriteriaType CriteriaType `json:"criteria_type"`
}

type CriteriaWeight struct {
	gorm.Model
	ID         int        `json:"id"`
	CriteriaID CriteriaID `json:"criteria_id,omitempty"`
	CategoryID int        `json:"category_id,omitempty"`
	Weight     float64    `json:"weight,omitempty"`
}

type CriteriaScore struct {
	CriteriaData   CriteriaData   `json:"criteria_data"`
	CriteriaWeight CriteriaWeight `json:"criteria_weight"`
	Score          int            `json:"score,omitempty"`
}

type Company struct {
	gorm.Model
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Ticker     string `json:"ticker,omitempty"`
	CategoryID int    `json:"category_id,omitempty"`
}

type Category struct {
	gorm.Model
	ID              int              `json:"id,omitempty"`
	Name            string           `json:"name,omitempty"`
	CriteriaWeights []CriteriaWeight `json:"criteria_weights,omitempty"`
}

type ESGRatingResult struct {
	Rating         float64         `json:"rating"`
	ERating        float64         `json:"e_rating"`
	SRating        float64         `json:"s_rating"`
	GRating        float64         `json:"g_rating"`
	CriteriaScores []CriteriaScore `json:"criteria_scores"`
}

type Report struct {
	ESGRatingResult ESGRatingResult `json:"esg_rating_result"`
	Company         Company         `json:"company"`
}

type Visualization struct {
	ESGRatingResult ESGRatingResult `json:"esg_rating_result"`
	Company         Company         `json:"company"`
}
