package service

import (
	"database/sql"
	"fmt"
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

// UpdateAllMoniestsPostCryptoStatistics updates post crypto statistics for all moniests
func (service *Service) UpdateAllMoniestsPostCryptoStatistics(ctx *gin.Context) error {
	// STEP: 7 days
	err := service.Store.UpdateAllMoniestsPostCryptoStatistics_7days(ctx)
	if err != nil {
		return fmt.Errorf("error while updating all moniests post crypto statistics [7 days]: %s", err.Error())
	}

	// STEP: 30 days
	err = service.Store.UpdateAllMoniestsPostCryptoStatistics_30days(ctx)
	if err != nil {
		return fmt.Errorf("error while updating all moniests post crypto statistics [30 days]: %s", err.Error())
	}

	// STEP: total
	err = service.Store.UpdateAllMoniestsPostCryptoStatistics_total(ctx)
	if err != nil {
		return fmt.Errorf("error while updating all moniests post crypto statistics [total]: %s", err.Error())
	}

	return nil
}

// UpdateMoniestsPostCryptoStatistics updates post crypto statistics for specific moniests
func (service *Service) UpdateMoniestsPostCryptoStatistics(ctx *gin.Context, moniest_ids []string) error {
	if len(moniest_ids) == 0 {
		return nil
	}

	// 7 days
	err := service.Store.UpdateMoniestsPostCryptoStatistics_7days(ctx, moniest_ids)
	if err != nil {
		return fmt.Errorf("error while updating moniests post crypto statistics [7 days]: %s, moniestIDs: %+q", err.Error(), moniest_ids)
	}

	// STEP: 30 days
	err = service.Store.UpdateMoniestsPostCryptoStatistics_30days(ctx, moniest_ids)
	if err != nil {
		return fmt.Errorf("error while updating moniests post crypto statistics [30 days]: %s, moniestIDs: %+q", err.Error(), moniest_ids)
	}

	// STEP: total
	err = service.Store.UpdateMoniestsPostCryptoStatistics_total(ctx, moniest_ids)
	if err != nil {
		return fmt.Errorf("error while updating moniests post crypto statistics [total]: %s, moniestIDs: %+q", err.Error(), moniest_ids)
	}

	return nil
}
