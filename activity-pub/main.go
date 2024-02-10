package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler)
	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe(":3000", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from activity-pub!")
}
