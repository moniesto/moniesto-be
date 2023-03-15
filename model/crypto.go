package model

type GetCurrenciesRequest struct {
	Name string `form:"name" json:"name" binding:"required,min=1"`
}

type Currency struct {
	Currency string `json:"currency"`
	Price    string `json:"price"`
}

type GetCurrenciesAPIResponse []struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type GetCurrencyAPIResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type GetHistoryDataAPIResponse [][]interface{}

type History struct {
	OpenTime   int64
	OpenPrice  string
	HighPrice  string
	LowPrice   string
	ClosePrice string
	Volume     string
	CloseTime  int64
}
