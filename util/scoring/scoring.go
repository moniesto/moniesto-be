package scoring

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/systemError"
)

var calculateApproximateURI = "/calculateApproximateScore"
var calculateURI = "/calculateScore"

// CalculateApproxScore calculates the approximately score of the post
func CalculateApproxScore(endDate time.Time, startPrice float64, endPrice float64, direction string, config config.Config) float64 {
	client := resty.New()

	var requestBody model.CalculateApproximateScoreRequest = model.CalculateApproximateScoreRequest{
		StartDate:  util.DateToTimestamp(util.Now()),
		EndDate:    util.DateToTimestamp(endDate),
		StartPrice: startPrice,
		EndPrice:   endPrice,
		Direction:  direction,
	}

	var response model.CalculateApproximateScoreResponse

	api_link := config.ScoringServiceURL + calculateApproximateURI

	resp, err := client.R().SetResult(&response).SetBody(requestBody).Post(api_link)

	if err != nil || resp.StatusCode() >= 400 {
		systemError.Log("calculate approximate score error", err.Error())
		return -1
	}

	return response.Score
}

// CalculateScore calculates the exact score of the post (run only for active posts)
func CalculateScore(requestBody model.CalculateScoreRequest, config config.Config) (model.CalculateScoreResponse, error) {

	client := resty.New()

	var response model.CalculateScoreResponse

	api_link := config.ScoringServiceURL + calculateURI

	resp, err := client.R().SetResult(&response).SetBody(requestBody).Post(api_link)

	if err != nil || resp.StatusCode() >= 400 {
		systemError.Log("calculate score error", err.Error())
		return model.CalculateScoreResponse{}, fmt.Errorf("calculate score error")
	}

	return response, nil
}
