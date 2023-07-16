package binance

import (
	"encoding/json"
	"fmt"

	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/core"
)

func CreatePayout(config config.Config, amount, operationFeePercentage float64, moniestPayoutType, moniestPayoutValue string) (string, PayoutTransferResponse, error) {

	uri := BASE_URL + CREATE_PAYOUT_PATH

	request_id := core.CreateID()
	merchant_send_id := core.CreateID()
	updatedAmount := core.GetAmountAfterCommission(amount, operationFeePercentage)

	body := CreatePayoutRequest{
		RequestID:   request_id,
		BatchName:   "Batch-" + request_id,
		Currency:    "USDT",
		TotalAmount: updatedAmount,
		TotalNumber: 1,
		BizScene:    "MERCHANT_PAYMENT",
		TransferDetailList: []TransferDetail{
			{
				MerchantSendId: merchant_send_id,
				TransferAmount: updatedAmount,
				ReceiveType:    moniestPayoutType,
				TransferMethod: "FUNDING_WALLET",
				Receiver:       moniestPayoutValue,
			},
		},
	}

	req, err := requestWithBinanceHeader(body, config)
	if err != nil {
		return request_id, PayoutTransferResponse{}, err
	}

	resp, err := req.SetBody(body).Post(uri)
	if err != nil {
		return request_id, PayoutTransferResponse{}, fmt.Errorf("error while sending request")
	}

	responseBody := CreatePayoutResponse{}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		return request_id, PayoutTransferResponse{}, fmt.Errorf("error while marshall response body")
	}

	if responseBody.Status == BINANCE_REQUEST_STATUS_FAIL {
		return request_id, PayoutTransferResponse{}, fmt.Errorf("error while creating payout: %s", responseBody.ErrorMessage)
	}

	if responseBody.Status == BINANCE_REQUEST_STATUS_SUCCESS {
		return request_id, responseBody.Data, nil
	}

	return request_id, PayoutTransferResponse{}, fmt.Errorf("[unkwown status] error while creating payout: %s", responseBody.ErrorMessage)
}
