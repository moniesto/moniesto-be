package binance

type RequestHeader struct {
	ContentType             string
	BinancepayTimestamp     int64
	BinancepayNonce         string
	BinancepayCertificateSN string
	BinancepaySignature     string
}

type CreateOrderRequest struct {
	Env             Env     `json:"env"`
	MerchantTradeNo string  `json:"merchantTradeNo"`
	OrderAmount     float64 `json:"orderAmount"`
	Currency        string  `json:"currency"` // BUSD | USDT | MBOX
	Goods           Goods   `json:"goods"`
	// Shipping        Shipping `json:"Shipping"`
}

type Env struct {
	TerminalType string `json:"terminalType"` // APP | WEB | WAP | MINI_PROGRAM | OTHERS
}

type Goods struct {
	GoodsType        string `json:"goodsType"`     // 01 | 02
	GoodsCategory    string `json:"goodsCategory"` // 0000: Electronics & Computers
	ReferenceGoodsId string `json:"referenceGoodsId"`
	GoodsName        string `json:"goodsName"`
	// GoodsUnitAmount  GoodsUnitAmount `json:"goodsUnitAmount"`
}

type GoodsUnitAmount struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type Shipping struct {
	ShippingName    ShippingName    `json:"shippingName"`
	ShippingAddress ShippingAddress `json:"shippingAddress"`
}

type ShippingName struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type ShippingAddress struct {
	Region string `json:"region"`
}
