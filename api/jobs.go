package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util"
	"github.com/moniesto/moniesto-be/util/system"
)

func (server *Server) Analyzer() {
	system.Log("JOB TRIGGER: Update Post Status")
	defer system.Timer("Analyzer")()

	ctx := gin.Context{}

	// STEP: get active posts
	activePosts, err := server.service.GetAllActivePosts()
	if err != nil {
		system.LogError("JOB ERROR: POST STATUS => db error while getting active posts")
		return
	}

	system.Log("# Active posts", len(activePosts))

	moniestIDs := []string{}

	// STEP: update active posts' status
	for i, post := range activePosts {
		system.Log(fmt.Sprintf("post start: %d id: %s\n", i, post.ID))

		status, err := server.service.UpdatePostStatus(post)
		if err != nil {
			system.LogError(fmt.Sprintf("JOB ERROR: POST STATUS => %s", err.Error()))
		}

		if status == db.PostCryptoStatusFail || status == db.PostCryptoStatusSuccess {
			if !util.Contains(moniestIDs, post.MoniestID) {
				moniestIDs = append(moniestIDs, post.MoniestID)
			}
		}
	}

	// STEP: update moniests' post crypto statistics that are finished [fail | success]
	err = server.service.UpdateMoniestsPostCryptoStatistics(&ctx, moniestIDs)
	if err != nil {
		system.LogError("JOB ERROR: UPDATE MONIEST POST STATISTICS", err.Error())
	}
}

func (server *Server) UpdateMoniestPostCryptoStatistics() {
	system.Log("JOB TRIGGER: Update Moniest Stats")
	defer system.Timer("Updating All Moniests Post Crypto Statistics")()

	ctx := gin.Context{}

	err := server.service.UpdateAllMoniestsPostCryptoStatistics(&ctx)
	if err != nil {
		system.LogError("JOB ERROR: UPDATE ALL MONIESTS POST STATISTICS", err.Error())
	}
}

func (server *Server) PayoutToMoniest() {
	system.Log("JOB TRIGGER: Payout To Moniest")
	defer system.Timer("Payout to Moniest")()

	ctx := gin.Context{}

	pendingPayouts, err := server.service.GetAllPendingPayouts(&ctx)
	if err != nil {
		system.LogError("JOB ERROR: PAYOUT => db error while getting pending payouts", err.Error())
		return
	}

	for _, pendingPayout := range pendingPayouts {
		err := server.service.PayoutToMoniest(&ctx, pendingPayout)
		if err != nil {
			system.LogError(fmt.Sprintf("JOB ERROR: PAYOUT => %s", err.Error()))
		}
	}
}

func (server *Server) DetectExpiredActiveSubscriptions() {
	system.Log("JOB TRIGGER: Detect Expired Active Subscriptions")
	defer system.Timer("Detect Expired Active Subscriptions")()

	ctx := gin.Context{}

	expiredSubscriptions, err := server.service.GetExpiredActiveSubscriptions(&ctx)
	if err != nil {
		system.LogError(fmt.Sprintf("JOB ERROR: EXPIRED SUBSCRIPTIONS => %s", err.Error()))
		return
	}

	system.Log("# of expired subscriptions", len(expiredSubscriptions))

	for _, expiredSubscription := range expiredSubscriptions {
		err := server.service.DeactivateExpiredSubscriptions(&ctx, expiredSubscription)
		if err != nil {
			system.LogError(fmt.Sprintf("JOB ERROR: EXPIRED SUBSCRIPTIONS => %s, user subsription ID: %s", err, expiredSubscription.ID))
		}
	}
}

func (server *Server) DetectExpiredPendingTransaction() {
	system.Log("JOB TRIGGER: Detect Expired Pending Transaction")
	defer system.Timer("Detect Expired Pending Transaction")()

	ctx := context.Background()

	expiredPendingTransactions, err := server.service.GetExpiredPendingBinanceTransactions(ctx)
	if err != nil {
		system.LogError(fmt.Sprintf("JOB ERROR: EXPIRED PENDING BINANCE TRANSACTIONS => %s", err.Error()))
	}

	system.Log("# of expired pending binance transactions", len(expiredPendingTransactions))

	for _, expiredPendingTransaction := range expiredPendingTransactions {
		err = server.service.UpdateExpiredPendingBinanceTransaction(ctx, expiredPendingTransaction.ID)
		if err != nil {
			system.LogError(fmt.Sprintf("JOB ERROR: EXPIRED PENDING BINANCE TRANSACTIONS => %s", err.Error()))
		}
	}
}

func (server *Server) MoniestRobot() {
	err := server.service.CreateRandomPost()
	if err != nil {
		system.LogError(err.Error())
	}
}
