package util

import (
	"log"
	"strconv"
)

func ToInt(value string) int {
	integer, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Could not convert \"%s\" into an int", value)
		panic(err)
	}
	return integer
}
