package util

import (
	"errors"
	"strconv"
	"time"
)

const invalidDuration = -1

func AtoTime(a string) (time.Duration, error) {
	if a == "" {
		return invalidDuration, errors.New("Please input time")
	}
	num, err := strconv.Atoi(a)
	if err != nil {
		return invalidDuration, err
	}
	d := time.Duration(num) * time.Minute
	return d, err
}
