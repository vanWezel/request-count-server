package helper

import (
	"os"
	"strconv"
)

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetenvToi(key, fallback string) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = fallback
	}

	integer, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return integer
}
