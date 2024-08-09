package main

import (
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello ProductStar!!! You are the best!!!"))
}

func main() {
	http.HandleFunc("/", myHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panicf("error while serving: %s", err)
	}
}

