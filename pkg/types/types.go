package types

type CriteriaData struct {
	CriteriaID int
	Name       string
	Weight     CriteriaWeight
}

type CriteriaWeight struct {
	CriteriaID int
	CategoryID int
	Weight     float64
}

type CriteriaScore struct {
	CriteriaData   CriteriaData
	CriteriaWeight CriteriaWeight
	Score          int
}

type Company struct {
	ID         int
	Name       string
	Ticker     string
	CategoryID int
}

type Category struct {
	ID              int
	Name            string
	CriteriaWeights []CriteriaWeight
}

type ESGRatingResult struct {
	Rating  float64
	ERating float64
	SRating float64
	GRating float64
}

type Report struct{}

type Visualization struct{}
