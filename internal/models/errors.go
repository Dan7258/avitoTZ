package models

import "errors"

var (
	ZeroUpdatedRowsError = errors.New("zero updated rows")
	UserNotFoundError    = errors.New("user not found")
)
