package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /.well-known/webfinger", GETWebfinger)
	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe(":3000", mux)
}

func GETWebfinger(w http.ResponseWriter, r *http.Request) {
	resource := r.URL.Query().Get("resource")
	fmt.Fprintf(w, "Hello from webfinger! You requested resource: %s", resource)
}
