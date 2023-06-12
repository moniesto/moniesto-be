package binance

import "time"

const CREATE_ORDER_STATUS_SUCCESS = "SUCCESS"
const CREATE_ORDER_STATUS_FAIL = "FAIL"
const ORDER_EXPIRE_TIME = time.Minute * 5 // 5 mins

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
	ReturnURL       string  `json:"returnUrl"`
	CancelURL       string  `json:"cancelUrl"`
	WebhookURL      string  `json:"webhookUrl"`
	OrderExpireTime int64   `json:"orderExpireTime"`
}

type CreateOrderResponse struct {
	Status string    `json:"status"`
	Code   string    `json:"code"`
	Data   OrderData `json:"data"`

	// only failure | error case
	ErrorMessage string `json:"errorMessage"`
}

type OrderData struct {
	Currency     string `json:"currency"`
	TotalFee     string `json:"totalFee"`
	PrepayId     string `json:"prepayId"`
	TerminalType string `json:"terminalType"`
	ExpireTime   int    `json:"expireTime"`

	// pay links
	QrcodeLink   string `json:"qrcodeLink"`
	QrContent    string `json:"qrContent"`
	CheckoutUrl  string `json:"checkoutUrl"`
	Deeplink     string `json:"deeplink"`
	UniversalUrl string `json:"universalUrl"`
}

type Env struct {
	TerminalType string `json:"terminalType"` // APP | WEB | WAP | MINI_PROGRAM | OTHERS
}

type Goods struct {
	GoodsType        string `json:"goodsType"`     // 01 | 02
	GoodsCategory    string `json:"goodsCategory"` // 0000: Electronics & Computers
	ReferenceGoodsId string `json:"referenceGoodsId"`
	GoodsName        string `json:"goodsName"`
}
