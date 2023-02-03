package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	deadline := time.Now().Add(10 * time.Minute)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if time.Now().After(deadline) {
			// Sleep for 3 seconds to trigger health check timeout
			fmt.Printf("Sleeping for 3 seconds to trigger health check timeout\n")
			time.Sleep(3 * time.Second)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Printf("endpoint called\n")
		fmt.Fprintf(w, "hello\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
