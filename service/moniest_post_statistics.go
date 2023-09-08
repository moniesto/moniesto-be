package service

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/core"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/clientError"
)

func (service *Service) CreateMoniestPostCryptoStatistics(ctx *gin.Context, moniest_id string) (db.MoniestPostCryptoStatistic, error) {

	params := db.CreateMoniestPostCryptoStatisticsParams{
		ID:            core.CreateID(),
		MoniestID:     moniest_id,
		Pnl7days:      sql.NullFloat64{Float64: 0, Valid: true},
		Roi7days:      sql.NullFloat64{Float64: 0, Valid: true},
		WinRate7days:  sql.NullFloat64{Float64: 0, Valid: true},
		Posts7days:    []string{},
		Pnl30days:     sql.NullFloat64{Float64: 0, Valid: true},
		Roi30days:     sql.NullFloat64{Float64: 0, Valid: true},
		WinRate30days: sql.NullFloat64{Float64: 0, Valid: true},
		Posts30days:   []string{},
		PnlTotal:      sql.NullFloat64{Float64: 0, Valid: true},
		RoiTotal:      sql.NullFloat64{Float64: 0, Valid: true},
		WinRateTotal:  sql.NullFloat64{Float64: 0, Valid: true},
	}

	postStatistics, err := service.Store.CreateMoniestPostCryptoStatistics(ctx, params)
	if err != nil {
		return db.MoniestPostCryptoStatistic{}, clientError.CreateError(http.StatusInternalServerError, clientError.Moniest_CreatePostStatistics_ServerErrorOnCreate)
	}

	return postStatistics, nil
}
