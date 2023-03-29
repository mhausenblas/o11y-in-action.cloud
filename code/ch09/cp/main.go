package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	_ "net/http/pprof" // <1>
)

func quickTask() (result string) { // <2>
	i := 600000
	for i > 0 {
		h := sha256.New()
		h.Write([]byte("just some busy work string hashing"))
		result = fmt.Sprintf("%x", h.Sum(nil))
		i--
	}
	return result
}

func slowTask() (result string) { // <3>
	i := 5400000
	for i > 0 {
		h := sha256.New()
		h.Write([]byte("just some busy work string hashing"))
		result = fmt.Sprintf("%x", h.Sum(nil))
		i--
	}
	return result
}

func main() {
	go func() { // <4>
		http.ListenAndServe("localhost:6060", nil)
	}()

	for true {
		result := quickTask()
		fmt.Printf("Quick task result is %s\n", result)
		result = slowTask()
		fmt.Printf("Slow task result is %s\n", result)
	}
}
