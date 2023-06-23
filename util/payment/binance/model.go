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
	ReturnURL       string  `json:"returnUrl"`
	CancelURL       string  `json:"cancelUrl"`
	WebhookURL      string  `json:"webhookUrl"`
	OrderExpireTime int64   `json:"orderExpireTime"`
}

type QueryOrderRequest struct {
	MerchantTradeNo string `json:"merchantTradeNo"`
}

type QueryOrderResponse struct {
	Status string         `json:"status"`
	Code   string         `json:"code"`
	Data   QueryOrderData `json:"data"`

	// only failure | error case
	ErrorMessage string `json:"errorMessage"`
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

type QueryOrderData struct {
	MerchantID      int    `json:"merchantId"`
	PrepayID        string `json:"prepayId"`
	MerchantTradeNo string `json:"merchantTradeNo"`
	Status          string `json:"status"`
	Currency        string `json:"currency"`
	CreateTime      int64  `json:"createTime"`
	OrderAmount     string `json:"orderAmount"`

	// paid specific
	OpenUserId   string           `json:"openUserId"`
	TransactTime int64            `json:"transactTime"`
	PaymentInfo  QueryPaymentInfo `json:"paymentInfo"`
}

type QueryPaymentInfo struct {
	PayerId             string                     `json:"payerId"`
	PayMethod           string                     `json:"payMethod"`
	PaymentInstructions []QueryPaymentInstructions `json:"paymentInstructions"`
	Channel             string                     `json:"channel"`
}

type QueryPaymentInstructions struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
	Price    string `json:"price"`
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

type WebhookRequest struct {
	BizType     string      `json:"bizType"`
	BizStatus   string      `json:"bizStatus"`
	WebhookData WebhookData `json:"data"`
	PayerId     int64       `json:"payerId"`
}

type WebhookData struct {
	MerchantTradeNo string           `json:"merchantTradeNo"`
	PaymentInfo     QueryPaymentInfo `json:"paymentInfo"`
}

/*
- Example webhook data
	{
		"bizType": "PAY",
		"data": {
			"merchantTradeNo": "755f2a5bdc42444991b08124eda15638",
			"productType": "02",
			"productName": "Moniest 1 - A",
			"transactTime": 1685921049737,
			"tradeType": "WEB",
			"totalFee": 1e-7,
			"currency": "USDT",
			"transactionId": "P_A1BQS87BCQ171112",
			"commission": 0,
			"paymentInfo": {
			"payerId": 741232235,
			"payMethod": "funding",
			"paymentInstructions": [
				{
				"currency": "USDT",
				"amount": 1e-7,
				"price": 1
				}
			],
			"channel": "DEFAULT"
			}
		},
		"bizIdStr": "232103202548367360",
		"bizId": 232103202548367360,
		"bizStatus": "PAY_SUCCESS"
		}
*/
