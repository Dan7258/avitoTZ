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
	errorResp := ErrorResponse{Code: status, Message: msg}
	resp := createResponse(errorResp, "error")
	w.WriteHeader(httpCode)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(resp)
	w.Header().Set("Content-Type", "application/json")
}

func jsonOK(w http.ResponseWriter, data interface{}, keyName string, httpCode int) {
	resp := createResponse(data, keyName)
	w.WriteHeader(httpCode)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(resp)
	w.Header().Set("Content-Type", "application/json")
}

func createResponse(data interface{}, key string) map[string]interface{} {
	return map[string]interface{}{
		key: data,
	}
}
