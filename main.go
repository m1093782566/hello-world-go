package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "Failed to get hostname, err: %v\n", err)
	} else {
		fmt.Fprintf(w, "Hi there, I am in %s!\n", hostName)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
