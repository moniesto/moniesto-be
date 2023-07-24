package api

import (
	"context"
	"fmt"

	"github.com/moniesto/moniesto-be/util/systemError"
)

func (server *Server) UpdatePostStatus() {
	fmt.Println("JOB TRIGGER: Update Post Status")

	activePosts, err := server.service.GetAllActivePosts()
	if err != nil {
		systemError.Log("JOB ERROR: POST STATUS => db error while getting active posts")
		return
	}

	fmt.Println("# Active posts", len(activePosts))

	for i, post := range activePosts {
		fmt.Printf("post start: %d id: %s\n", i, post.ID)
		err := server.service.UpdatePostStatus(post)
		if err != nil {
			systemError.Log(fmt.Sprintf("JOB ERROR: POST STATUS => %s", err))
		}
	}
}

func (server *Server) PayoutToMoniest() {
	fmt.Println("JOB TRIGGER: Payout To Moniest")

	pendingPayouts, err := server.service.GetAllPendingPayouts()
	if err != nil {
		systemError.Log("JOB ERROR: PAYOUT => db error while getting pending payouts", err.Error())
		return
	}

	for _, pendingPayout := range pendingPayouts {
		err := server.service.PayoutToMoniest(pendingPayout)
		if err != nil {
			systemError.Log(fmt.Sprintf("JOB ERROR: PAYOUT => %s", err))
		}
	}
}

func (server *Server) DetectExpiredActiveSubscriptions() {
	fmt.Println("JOB TRIGGER: Detect Expired Active Subscriptions")

	ctx := context.Background()

	expiredSubscriptions, err := server.service.GetExpiredActiveSubscriptions(ctx)
	if err != nil {
		systemError.Log(fmt.Sprintf("JOB ERROR: EXPIRED SUBSCRIPTIONS => %s", err))
		return
	}

	fmt.Println("# of expired subscriptions", len(expiredSubscriptions))

	for _, expiredSubscription := range expiredSubscriptions {
		err := server.service.DeactivateExpiredSubscriptions(ctx, expiredSubscription)
		if err != nil {
			systemError.Log(fmt.Sprintf("JOB ERROR: EXPIRED SUBSCRIPTIONS => %s, user subsription ID: %s", err, expiredSubscription.ID))
		}
	}

	// TODO: send email to expired subscriptions [users]
}

func (server *Server) DetectExpiredPendingTransaction() {
	fmt.Println("JOB TRIGGER: Detect Expired Pending Transaction")

	ctx := context.Background()

	expiredPendingTransactions, err := server.service.GetExpiredPendingBinanceTransactions(ctx)
	if err != nil {
		systemError.Log(fmt.Sprintf("JOB ERROR: EXPIRED PENDING BINANCE TRANSACTIONS => %s", err))
	}

	fmt.Println("# of expired pending binance transactions", len(expiredPendingTransactions))

	for _, expiredPendingTransaction := range expiredPendingTransactions {
		err = server.service.UpdateExpiredPendingBinanceTransaction(ctx, expiredPendingTransaction.ID)
		if err != nil {
			systemError.Log(fmt.Sprintf("JOB ERROR: EXPIRED PENDING BINANCE TRANSACTIONS => %s", err))
		}
	}

}
