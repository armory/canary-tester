package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

func main() {
	fmt.Println("Canary Tester started.")
	failureRate := 0.10
	setRootHandlers()
	setToggleHandlers(&failureRate)
	go dataDogPulse(&failureRate)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Done.")
}

func setRootHandlers() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received", r)
		w.WriteHeader(200)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received", r)
		w.Write([]byte(`{"status":"ok"`))
	})
}

func setToggleHandlers(failureRate *float64) {
	http.HandleFunc("/fail/true", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Emitting more failures.")
		fmt.Println("Received", r)
		*failureRate = 0.90
	})
	http.HandleFunc("/fail/false", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Emitting more successes.")
		fmt.Println("Received", r)
		*failureRate = 0.10
	})
}

func dataDogPulse(failureRate *float64) {
	c, _ := statsd.New("127.0.0.1:8125")
	c.Namespace = "canarytester."
	for {
		r := rand.Float64()
		name := "success"
		if r < *failureRate {
			name = "failure"
		}
		fmt.Println(name)
		err := c.Incr(name, nil, 1.0)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
