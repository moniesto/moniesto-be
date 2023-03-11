package model

type CalculateApproximateScoreRequest struct {
	StartDate  int64   `json:"startDate"`
	EndDate    int64   `json:"endDate"`
	StartPrice float64 `json:"startPrice"`
	EndPrice   float64 `json:"endPrice"`
	Direction  string  `json:"direction"`
}

type CalculateApproximateScoreResponse struct {
	Score float64 `json:"score"`
}
