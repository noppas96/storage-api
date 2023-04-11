package main

import "net/http"

func main() {
	http.HandleFunc("/health", HealthCheck)
	http.ListenAndServe(":3010", nil)
}
