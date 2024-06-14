package iid

import (
	"errors"
	"strconv"
)

var (
	ErrTooManyTokens = errors.New("batch import only supports importing " + strconv.Itoa(maxTokens) + " at a time")
)
