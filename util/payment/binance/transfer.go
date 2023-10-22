package binance

import (
	"encoding/json"
	"fmt"

	"github.com/moniesto/moniesto-be/config"
	"github.com/moniesto/moniesto-be/core"
)

func CreateTransfer(config config.Config, amount, operationFeePercentage float64, transferType, receiveType, receiveValue, remark string) (CreateTransferRequest, CreateTransferResponse, TransferResponse, error) {
	uri := BASE_URL + CREATE_PAYOUT_PATH

	request_id := core.CreatePlainID()
	merchant_send_id := core.CreatePlainID()
	updatedAmount := core.GetAmountAfterCommission(amount, operationFeePercentage, &config)

	body := CreateTransferRequest{
		RequestID:   request_id,
		BatchName:   "Batch-" + request_id[0:26],
		Currency:    "USDT",
		TotalAmount: updatedAmount,
		TotalNumber: 1,
		BizScene:    transferType,
		TransferDetailList: []TransferDetail{
			{
				MerchantSendId: merchant_send_id,
				TransferAmount: updatedAmount,
				ReceiveType:    receiveType,
				TransferMethod: "FUNDING_WALLET",
				Receiver:       receiveValue,
				Remark:         remark,
			},
		},
	}

	req, err := requestWithBinanceHeader(body, config)
	if err != nil {
		return body, CreateTransferResponse{}, TransferResponse{}, err
	}

	resp, err := req.SetBody(body).Post(uri)
	if err != nil {
		return body, CreateTransferResponse{}, TransferResponse{}, fmt.Errorf("error while sending request")
	}

	responseBody := CreateTransferResponse{}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		return body, responseBody, TransferResponse{}, fmt.Errorf("error while marshall response body")
	}

	if responseBody.Status == BINANCE_REQUEST_STATUS_FAIL {
		return body, responseBody, TransferResponse{}, fmt.Errorf("error while creating payout: %s", responseBody.ErrorMessage)
	}

	if responseBody.Status == BINANCE_REQUEST_STATUS_SUCCESS {
		return body, responseBody, responseBody.Data, nil
	}

	return body, responseBody, TransferResponse{}, fmt.Errorf("[unkwown status] error while creating payout: %s", responseBody.ErrorMessage)
}
