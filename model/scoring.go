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

type CalculateScoreRequest struct {
	Parity               string      `json:"parity"`
	StartPrice           float64     `json:"startPrice"`
	StartDate            int64       `json:"startDate"`
	EndDate              int64       `json:"endDate"`
	Target1              float64     `json:"target1"`
	Target2              float64     `json:"target2"`
	Target3              float64     `json:"target3"`
	Stop                 float64     `json:"stop"`
	Direction            string      `json:"direction"`
	LastCronJobTimeStamp interface{} `json:"lastCronJobTimeStamp"`
	LastTargetHit        float64     `json:"lastTargetHit"`
}

type CalculateScoreResponse struct {
	Finished             bool        `json:"finished"`
	Score                float64     `json:"score"`
	Success              bool        `json:"success"`
	LastTargetHit        float64     `json:"lastTargetHit"`
	LastCronJobTimeStamp interface{} `json:"lastCronJobTimeStamp"`
}
