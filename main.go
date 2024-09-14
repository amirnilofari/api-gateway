package main

import (
	"fmt"
	"io"
	"net/http"
)

var routes = map[string]string{
	"/users":  "http://localhost:8081",
	"/orders": "http://localhost:8082",
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", routeHandler)

	fmt.Println("API Gateway running on port 8080")
	http.ListenAndServe(":8080", mux)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	backendURL, ok := routes[r.URL.Path]
	if !ok {
		http.Error(w, "Route not found", http.StatusNotFound)
		return
	}

	resp, err := http.Get(backendURL + r.URL.Path)
	if err != nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
