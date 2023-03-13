package Helper

import (
	"net/http"
)

type WrapResponseWriter struct {
	http.ResponseWriter
	ResponseBody []byte
	HttpCode     int
}

func NewWrapResponseWriter(w http.ResponseWriter) *WrapResponseWriter {
	return &WrapResponseWriter{ResponseWriter: w}
}

// Write body request and get size and data
func (w *WrapResponseWriter) Write(data []byte) (int, error) {
	size, err := w.ResponseWriter.Write(data)
	w.ResponseBody = data
	return size, err
}

// Write Header and assign http code to this struct
func (w *WrapResponseWriter) WriteHeader(statusCode int) {
	w.HttpCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
