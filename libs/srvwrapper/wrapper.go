package srvwrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Validator interface {
	Validate() error
}

type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type Wrapper[Req Validator, Res any] struct {
	fn func(ctx context.Context, req Req) (Res, error)
}

func New[Req Validator, Res any](fn func(ctx context.Context, req Req) (Res, error)) *Wrapper[Req, Res] {
	return &Wrapper[Req, Res]{
		fn: fn,
	}
}

func (w *Wrapper[Req, Res]) ServeHTTP(resWriter http.ResponseWriter, httpReq *http.Request) {
	ctx := httpReq.Context()

	var request Req
	err := json.NewDecoder(httpReq.Body).Decode(&request)
	if err != nil {
		sendResponse(resWriter, http.StatusBadRequest, getErrorResponse("decoding JSON", err))
		return
	}

	err = request.Validate()
	if err != nil {
		sendResponse(resWriter, http.StatusBadRequest, getErrorResponse("validating request", err))
		return
	}

	response, err := w.fn(ctx, request)
	if err != nil {
		sendResponse(resWriter, http.StatusInternalServerError, getErrorResponse("running handler", err))
		return
	}

	rawJSON, _ := json.Marshal(response)
	sendResponse(resWriter, http.StatusOK, rawJSON)
}

func sendResponse(w http.ResponseWriter, status int, body []byte) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(body)
}

func getErrorResponse(text string, err error) []byte {
	newError := errorResponse{Message: text, Error: err.Error()}
	buff := new(bytes.Buffer)
	_ = json.NewEncoder(buff).Encode(newError)
	return buff.Bytes()
}
