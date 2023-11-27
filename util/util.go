package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/moniesto/moniesto-be/util/system"
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

// SafeSQLNullToTime converts sql object to time pointer
func SafeSQLNullToTime(sqlNull sql.NullTime) *time.Time {
	if !sqlNull.Valid {
		return nil
	}

	return &sqlNull.Time
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

func Contains[T comparable](slice []T, element T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}

	return false
}

func Remove[T comparable](slice []T, element T) []T {
	for i, v := range slice {
		if v == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// emailWithoutLocal returns email without local(part after + sign) part
func EmailWithoutLocal(email string) (string, error) {
	// Split the email address at the "@" symbol to separate the local part and domain part
	parts := strings.Split(email, "@")

	if len(parts) != 2 {
		return "", fmt.Errorf("invalid email address")
	}

	localPart := parts[0]
	domain := parts[1]

	// Find the position of the "+" symbol in the local part
	plusIndex := strings.Index(localPart, "+")

	if plusIndex != -1 {
		// Remove everything from the "+" symbol to the end of the local part
		localPart = localPart[:plusIndex]
	}

	// Reconstruct the email address
	cleanedEmail := localPart + "@" + domain

	return cleanedEmail, nil
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

func StructToJSON(data any) string {
	str, err := json.Marshal(data)
	if err != nil {
		system.LogError("struct to json error", err.Error())
		return fmt.Sprintf("struct to json error: %s", err.Error())
	}

	return string(str)
}

func IsNight() bool {
	currentTime := Now()
	hour := currentTime.Hour()

	// Check if the hour is between 11 PM (23) and 8 AM (8)
	return hour >= 23 || hour < 8
}

// SimplifyRandomPrices -> mainPrice=1.0001, randomPrice=1.00009999999 -> turns randomPrice -> 1.00009
func SimplifyRandomPrices(mainPrice, randomPrice float64) float64 {
	decimalPlacesA := numDecPlaces(mainPrice)

	maxDecimalPlaces := decimalPlacesA + 1

	simplifiedB := cutDecimal(randomPrice, maxDecimalPlaces)

	return simplifiedB
}

func numDecPlaces(v float64) int {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	i := strings.IndexByte(s, '.')
	if i > -1 {
		return len(s) - i - 1
	}
	return 0
}

func cutDecimal(A float64, B int) float64 {
	pow := math.Pow(10, float64(B))
	result := math.Floor(A*pow) / pow
	return result
}
