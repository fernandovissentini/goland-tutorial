package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, r *http.Request) {
		globalGet(responseWriter, r)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		goodbyePath()
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}
}

func goodbyePath() {
	log.Println("Goodbye World")
}

func globalGet(responseWriter http.ResponseWriter, r *http.Request) {
	log.Println("Hello World")
	data, err1 := io.ReadAll(r.Body)
	if err1 != nil || len(data) < 1 {
		http.Error(responseWriter, "Error, no body", http.StatusBadRequest)
		return
	}
	log.Printf("Data %s", data)

	_, err := fmt.Fprintf(responseWriter, "Hello %s", data)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}
}
