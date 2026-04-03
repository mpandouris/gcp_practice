package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var (
	ServiceName = "basic-api"
)

func main() {
	port := flag.Int("port", 8080, "port for the server to listen on")
	flag.Parse()

	hostname, _ := os.Hostname()

	http.HandleFunc("/basic-api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "all good from %s\n", hostname)
	})

	http.HandleFunc("/basic-api/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome! This is a basic service to practice my skills :)\n")
	})

	fmt.Printf("Failed to start server: %v", http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
