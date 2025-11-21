package handler

import "avito/internal/models"

type ErrorResponse struct {
	Code    models.CodeError `json:"code"`
	Message string           `json:"message"`
}
