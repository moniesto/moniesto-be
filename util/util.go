package util

import (
	"database/sql"
	"math"
	"time"
)

const (
	MAX_LIMIT     = 50
	DEFAULT_LIMIT = 10

	DEFAULT_OFFSET   = 0
	MAX_SEARCH_LIMIT = 30
)

var (
	POST_FILTER_PNL        = "pnl"
	POST_FILTER_CREATED_AT = "created_at"
)

var PostSortingTypes []string = []string{
	POST_FILTER_CREATED_AT,
	POST_FILTER_PNL,
}

var defaultSortingType string = PostSortingTypes[1] // pnl

func SafeLimit(limit int) int {
	if limit > MAX_LIMIT {
		limit = MAX_LIMIT
	}

	if limit <= 0 {
		limit = DEFAULT_LIMIT
	}

	return limit
}

func SafeOffset(offset int) int {
	if offset < 0 {
		offset = DEFAULT_OFFSET
	}

	return offset
}

func SafePostSortBy(sortBy string) string {
	for _, v := range PostSortingTypes {
		if v == sortBy {
			return sortBy
		}
	}

	return defaultSortingType
}

func SafeSearchText(searchText string) string {

	if len(searchText) > MAX_SEARCH_LIMIT {
		searchText = searchText[0:MAX_SEARCH_LIMIT]
	}

	return searchText
}

// SafeFloat64ToSQLNull converts float pointer to sql object if valid
func SafeFloat64ToSQLNull(num *float64) sql.NullFloat64 {
	value := sql.NullFloat64{}

	if num != nil {
		value.Float64 = *num
		value.Valid = true
	}

	return value
}

// SafeSQLNullToFloat converts sql object to float pointer
func SafeSQLNullToFloat(sqlNull sql.NullFloat64) *float64 {
	if !sqlNull.Valid {
		return nil
	}

	return &sqlNull.Float64
}

func DateToTimestamp(date time.Time) int64 {
	return date.UnixNano() / int64(time.Millisecond)
}

func TimestampToDate(timestamp int64) time.Time {
	return time.Unix(timestamp/1000, 0)
}

func EarliestDate(date1 time.Time, date2 time.Time) time.Time {
	if date1.Before(date2) {
		return date1
	}
	return date2
}

func RoundAmountDown(fee float64) float64 {
	return (math.Floor(fee*100) / 100)
}

func RoundAmountUp(fee float64) float64 {
	return (math.Ceil(fee*100) / 100)
}

func Now() time.Time {
	return time.Now().UTC()
}
