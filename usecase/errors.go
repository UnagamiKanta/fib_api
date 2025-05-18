package usecase

import (
	"errors"
	"strconv"
)

var (
	ErrInvalidInput  = errors.New("n must be a non-negative integer")
	ErrTooLargeInput = errors.New("n is too large, please use less than " + strconv.Itoa(MAXINPUTNUM))
)
