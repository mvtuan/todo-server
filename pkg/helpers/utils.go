package helpers

import (
	"strconv"
)

func ParseInt64(s string, defaultNumber int64) int64 {
	if s == "" {
		return defaultNumber
	}

	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return defaultNumber
	}

	return num
}
func ParseInt(s string, defaultNumber int) int {
	if s == "" {
		return defaultNumber
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return defaultNumber
	}

	return num
}
