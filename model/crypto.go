package model

type GetCurrenciesRequest struct {
	Name       string `form:"name" json:"name" binding:"required,min=1"`
	MarketType string `form:"market_type" json:"market_type" binding:"required"`
}

type Currency struct {
	Currency string `json:"currency"`
	Price    string `json:"price"`
}

type GetCurrenciesAPIResponse []GetCurrencyAPIResponse

type GetCurrencyAPIResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type GetHistoryDataAPIResponse [][]interface{}

type History []any

type GetExchangeInfoResponse struct {
	Symbols []Symbol `json:"symbols"`
}

type Symbol struct {
	Symbol string `json:"symbol"`
	Status string `json:"status"` // TRADING | BREAK
}
