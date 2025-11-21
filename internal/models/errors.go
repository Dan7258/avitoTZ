package models

import "errors"

var (
	ZeroUpdatedRowsError = errors.New("zero updated rows")
	NotFoundError        = errors.New("not found")
	NotChangedError      = errors.New("not changed")
)
