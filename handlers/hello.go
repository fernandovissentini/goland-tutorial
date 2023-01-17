package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(responseWriter http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	data, err1 := io.ReadAll(r.Body)
	if err1 != nil || len(data) < 1 {
		http.Error(responseWriter, "Error, no body", http.StatusBadRequest)
		return
	}
	h.l.Printf("Data %s", data)

	_, err := fmt.Fprintf(responseWriter, "Hello %s", data)
	if err != nil {
		h.l.Printf("Error: %s", err.Error())
		return
	}
}
