package binance

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/core"
	"github.com/moniesto/moniesto-be/util"
)

func CreateOrder(ctx *gin.Context, config config.Config, transactionID string, amount float64, productName, returnURL, cancelURL, webhookURL string) (OrderData, error) {

	uri := BASE_URL + CREATE_ORDER_PATH

	body := CreateOrderRequest{
		Env: Env{
			TerminalType: "WEB",
		},
		MerchantTradeNo: transactionID,
		OrderAmount:     amount,
		Currency:        "USDT",
		Goods: Goods{
			GoodsType:        "02",
			GoodsCategory:    "0000",
			ReferenceGoodsId: core.CreatePlainID(),
			GoodsName:        productName,
		},
		WebhookURL:      webhookURL,
		ReturnURL:       returnURL,
		CancelURL:       cancelURL,
		OrderExpireTime: util.DateToTimestamp(util.Now().Add(ORDER_EXPIRE_TIME)),
	}

	req, err := requestWithBinanceHeader(body, config)
	if err != nil {
		return OrderData{}, err
	}

	resp, err := req.SetBody(body).Post(uri)
	if err != nil {
		return OrderData{}, fmt.Errorf("error while sending request")
	}

	responseBody := CreateOrderResponse{}

	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		return OrderData{}, fmt.Errorf("error while marshall response body")
	}

	if responseBody.Status == BINANCE_REQUEST_STATUS_FAIL {
		return OrderData{}, fmt.Errorf("error while creating order: %s", responseBody.ErrorMessage)
	}

	return responseBody.Data, nil
}
