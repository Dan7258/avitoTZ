package handler

import (
	"avito/internal/models"
)

type Handler struct {
	db models.Model
}

func InitHandler(model models.Model) *Handler {
	return &Handler{model}
}
