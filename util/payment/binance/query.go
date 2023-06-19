package binance

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/config"
)

func QueryOrder(ctx *gin.Context, config config.Config, transactionID string) (QueryOrderData, error) {
	uri := BASE_URL + QUERY_ORDER_PATH

	body := QueryOrderRequest{
		MerchantTradeNo: transactionID,
	}

	req, err := requestWithBinanceHeader(body, config)
	if err != nil {
		return QueryOrderData{}, err
	}

	resp, err := req.SetBody(body).Post(uri)
	if err != nil {
		return QueryOrderData{}, fmt.Errorf("error while sending request")
	}

	responseBody := QueryOrderResponse{}

	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		return QueryOrderData{}, fmt.Errorf("error while marshall response body")
	}

	if responseBody.Status == BINANCE_REQUEST_STATUS_FAIL {
		return QueryOrderData{}, fmt.Errorf("error while querying order: %s", responseBody.ErrorMessage)
	}

	return responseBody.Data, nil
}

/*
PENDING:

{
  "status": "SUCCESS",
  "code": "000000",
  "data": {
    "merchantId": 225467155,
    "prepayId": "234506691862945792",
    "merchantTradeNo": "bd1d51aee94a4f15a32e5673db6615bd",
    "status": "INITIAL",
    "currency": "USDT",
    "createTime": 1687037488960,
    "orderAmount": "1E-8"
  }
}
*/

/*
SUCCESS:

{
  "status": "SUCCESS",
  "code": "000000",
  "data": {
    "merchantId": 225467155,
    "prepayId": "234506691862945792",
    "transactionId": "M_S_234506915261751297",
    "merchantTradeNo": "bd1d51aee94a4f15a32e5673db6615bd",
    "status": "PAID",
    "currency": "USDT",
    "openUserId": "934620c827af6bb3d7ffc7f4e9126805",
    "transactTime": 1687037592784,
    "createTime": 1687037488960,
    "paymentInfo": {
      "payerId": "741232235",
      "payMethod": "spot",
      "paymentInstructions": [
        {
          "currency": "USDT",
          "amount": "0.000000010",
          "price": "1"
        }
      ],
      "channel": "DEFAULT"
    },
    "orderAmount": "1E-8"
  }
}
*/

/*
EXPIRED

{
  "status": "SUCCESS",
  "code": "000000",
  "data": {
    "merchantId": 225467155,
    "prepayId": "234508657160593408",
    "merchantTradeNo": "270c4807af134a3fb6d65441f05fba8e",
    "status": "EXPIRED",
    "currency": "USDT",
    "createTime": 1687038403632,
    "orderAmount": "1E-8"
  }
}
*/

/*
STATUS TYPES

"INITIAL", "PENDING", "PAID", "CANCELED", "ERROR", "REFUNDING", "REFUNDED", "EXPIRED"
*/

/*
ERROR

{
  "status": "FAIL",
  "code": "400202",
  "errorMessage": "Order not found."
}
*/
