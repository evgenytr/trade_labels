package domain

import "errors"

var (
	ErrorJSONTypeIncorrect      = errors.New("json type incorrect")
	ErrorRequestFormatIncorrect = errors.New("request format incorrect")
)
