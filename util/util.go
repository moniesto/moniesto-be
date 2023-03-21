package util

import "time"

const (
	MAX_LIMIT     = 50
	DEFAULT_LIMIT = 10

	DEFAULT_OFFSET   = 0
	MAX_SEARCH_LIMIT = 30
)

var (
	POST_FILTER_SCORE      = "score"
	POST_FILTER_CREATED_AT = "created_at"
)

var PostSortingTypes []string = []string{
	POST_FILTER_CREATED_AT,
	POST_FILTER_SCORE,
}

var defaultSortingType string = PostSortingTypes[1] // score

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

func DateToTimestamp(date time.Time) int64 {
	return date.UnixNano() / int64(time.Millisecond)
}
