package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Canary Tester started.")
	fail := shouldFail()
	http.HandleFunc("/fail/true", func(w http.ResponseWriter, r *http.Request) {
		fail = true
		fmt.Println("All responses from now on will be 200.")
		fmt.Println("Received", r)
	})
	http.HandleFunc("/fail/false", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received", r)
		fmt.Println("All responses from now on will be 500 errors.")
		fail = false
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received", r)
		if fail {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(200)
		}
	})
	http.ListenAndServe(":8080", nil)
	fmt.Println("Done.")
}

func shouldFail() bool {
	f := os.Getenv("RETURN_500")
	if f == "true" {
		return true
	}
	return false
}
