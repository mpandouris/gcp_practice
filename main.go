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

	hostname, _ = os.Hostname()
)

func main() {
	port := flag.Int("port", 8080, "port for the server to listen on")
	verbosity := flag.Int("verbosity", 0, "log verbosity - should be env driven")
	env := flag.String("env", "", "env/namespace of deployment")
	timeout := flag.Int("timeout", 10, "std timout for downstream http calls. default is 10s")
	flag.Parse()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "all good from %s\n", hostname)
	})

	http.HandleFunc("/basic-api/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome! This is a basic service to practice my skills :)\nArgs\n\tport: %d\n\tverbosity: %d\n\ttimeout: %ds\n\tenv: %s\n", *port, *verbosity, *timeout, *env)
	})

	fmt.Printf("Failed to start server: %v", http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
