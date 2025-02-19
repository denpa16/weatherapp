package http

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

func GetErrorResponse(w http.ResponseWriter, handlerName string, err error, statusCode int) {
	w.WriteHeader(statusCode)
	buf := bytes.NewBufferString(handlerName)
	buf.WriteString(": ")
	buf.WriteString(err.Error())
	buf.WriteString("\n")
	_, _ = w.Write(buf.Bytes())
}

func GetSuccessResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func GetSuccessResponseWithJSON(w http.ResponseWriter, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func GetSuccessCreatedResponseWithJSON(w http.ResponseWriter, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(body)
}

func GetSuccessResponseWithBytes(w http.ResponseWriter, body []byte, name string) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", name))
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}
