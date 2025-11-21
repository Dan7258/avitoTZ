package handler

import (
	"encoding/json"
	"net/http"
)

type CodeError string

const (
	TeamExists  CodeError = "TEAM_EXISTS"
	PrExists    CodeError = "PR_EXISTS"
	PrMerged    CodeError = "PR_MERGED"
	NotAssigned CodeError = "NOT_ASSIGNED"
	NoCandidate CodeError = "NO_CANDIDATE"
	NotFound    CodeError = "NOT_FOUND"
	NotValid    CodeError = "NOT_VALID"
)

type ErrorResponse struct {
	Code    CodeError `json:"code"`
	Message string    `json:"message"`
}

func jsonError(w http.ResponseWriter, status CodeError, msg string, httpCode int) {
	json.NewEncoder(w).Encode(&ErrorResponse{
		Code:    status,
		Message: msg,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
}
