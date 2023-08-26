package api

import (
	"context"
	"fmt"

	"github.com/moniesto/moniesto-be/util/system"
)

func (server *Server) UpdatePostStatus() {
	system.Log("JOB TRIGGER: Update Post Status")

	activePosts, err := server.service.GetAllActivePosts()
	if err != nil {
		system.LogError("JOB ERROR: POST STATUS => db error while getting active posts")
		return
	}

	system.Log("# Active posts", len(activePosts))

	for i, post := range activePosts {
		system.Log(fmt.Sprintf("post start: %d id: %s\n", i, post.ID))

		err := server.service.UpdatePostStatus(post)
		if err != nil {
			system.LogError(fmt.Sprintf("JOB ERROR: POST STATUS => %s", err))
		}
	}
}

func (server *Server) PayoutToMoniest() {
	system.Log("JOB TRIGGER: Payout To Moniest")

	pendingPayouts, err := server.service.GetAllPendingPayouts()
	if err != nil {
		system.LogError("JOB ERROR: PAYOUT => db error while getting pending payouts", err.Error())
		return
	}

	for _, pendingPayout := range pendingPayouts {
		err := server.service.PayoutToMoniest(pendingPayout)
		if err != nil {
			system.LogError(fmt.Sprintf("JOB ERROR: PAYOUT => %s", err))
		}
	}
}

func (server *Server) DetectExpiredActiveSubscriptions() {
	system.Log("JOB TRIGGER: Detect Expired Active Subscriptions")

	ctx := context.Background()

	expiredSubscriptions, err := server.service.GetExpiredActiveSubscriptions(ctx)
	if err != nil {
		system.LogError(fmt.Sprintf("JOB ERROR: EXPIRED SUBSCRIPTIONS => %s", err))
		return
	}

	system.Log("# of expired subscriptions", len(expiredSubscriptions))

	for _, expiredSubscription := range expiredSubscriptions {
		err := server.service.DeactivateExpiredSubscriptions(ctx, expiredSubscription)
		if err != nil {
			system.LogError(fmt.Sprintf("JOB ERROR: EXPIRED SUBSCRIPTIONS => %s, user subsription ID: %s", err, expiredSubscription.ID))
		}
	}

	// TODO: send email to expired subscriptions [users]
}

func (server *Server) DetectExpiredPendingTransaction() {
	system.Log("JOB TRIGGER: Detect Expired Pending Transaction")

	ctx := context.Background()

	expiredPendingTransactions, err := server.service.GetExpiredPendingBinanceTransactions(ctx)
	if err != nil {
		system.LogError(fmt.Sprintf("JOB ERROR: EXPIRED PENDING BINANCE TRANSACTIONS => %s", err))
	}

	system.Log("# of expired pending binance transactions", len(expiredPendingTransactions))

	for _, expiredPendingTransaction := range expiredPendingTransactions {
		err = server.service.UpdateExpiredPendingBinanceTransaction(ctx, expiredPendingTransaction.ID)
		if err != nil {
			system.LogError(fmt.Sprintf("JOB ERROR: EXPIRED PENDING BINANCE TRANSACTIONS => %s", err))
		}
	}

}
