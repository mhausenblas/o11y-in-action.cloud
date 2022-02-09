package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func handleEcho(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	log.WithFields(log.Fields{
		"service": "echo",
	}).Info("Got input ", message) //<2>
	if rand.Intn(100) > 90 {
		log.WithFields(log.Fields{
			"service": "echo",
		}).Error("Something really bad happened :(") //<3>
		w.WriteHeader(500)
		return
	}
	fmt.Fprintf(w, message)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFormatter(&log.JSONFormatter{}) // <1>
	http.HandleFunc("/echo", handleEcho)
	log.Fatal(http.ListenAndServe(":4242", nil))
}
