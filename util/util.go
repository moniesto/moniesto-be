package util

const (
	MAX_LIMIT     = 50
	DEFAULT_LIMIT = 10

	DEFAULT_OFFSET = 0
)

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
