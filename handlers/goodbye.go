package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(responseWriter http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye World")

	_, err := responseWriter.Write([]byte("Byeee"))
	if err != nil {
		g.l.Printf("Error: %s", err.Error())
		return
	}
}
