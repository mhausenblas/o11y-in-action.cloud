package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func handleEcho(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message") //<3>
	if rand.Intn(100) > 90 {                //<4>
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, message) //<5>
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/echo", handleEcho)         //<1>
	log.Fatal(http.ListenAndServe(":4242", nil)) //<2>
}
