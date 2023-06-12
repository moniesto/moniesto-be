package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/model"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/payment/binance"
	"github.com/moniesto/moniesto-be/util/systemError"
)

func (service *Service) CreateBinancePaymentTransaction(ctx *gin.Context, req model.SubscribeMoniestRequest, moniest db.GetMoniestByUsernameRow, userID string) {

	// STEP: create order in binance and get payment links
	product_name := getProductName(req, moniest)
	amount := getAmount(req, moniest)

	orderData, err := binance.CreateOrder(ctx, service.config, amount, product_name, req.ReturnURL, req.CancelURL)
	if err != nil {
		systemError.Log("create order error", err.Error())
		return
	}

	// STEP: add payment transactions to db

	fmt.Println("URL", orderData.UniversalUrl)
}

func getProductName(req model.SubscribeMoniestRequest, moniest db.GetMoniestByUsernameRow) string {
	return fmt.Sprintf("You are subscribing to %s %s at a monthly fee of $%.2f for %d months.", moniest.Name, moniest.Surname, util.RoundAmountDown(moniest.Fee), req.NumberOfMonths)
}

func getAmount(req model.SubscribeMoniestRequest, moniest db.GetMoniestByUsernameRow) float64 {
	return util.RoundAmountUp(float64(req.NumberOfMonths) * moniest.Fee)
}
