package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type Webfinger struct {
	Subject string   `json:"subject"`
	Aliases []string `json:"aliases"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/.well-known/webfinger", handleWebfinger)

	fmt.Println("Server is running on http://localhost:3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

func handleWebfinger(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	resource := r.URL.Query().Get("resource")
	username, host, err := parseResource(resource)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	webfinger, err := getWebfinger(username, host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(webfinger)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func parseResource(resource string) (username, host string, err error) {
	if len(resource) < 25 || !strings.HasPrefix(resource, "acct:") {
		return "", "", fmt.Errorf("invalid resource")
	}

	parts := strings.Split(resource[5:], "@")
	if len(parts) != 3 {
		return "", "", fmt.Errorf("invalid resource")
	}

	return parts[1], parts[2], nil
}

func getWebfinger(username, host string) (*Webfinger, error) {
	if host != "social-network.local" {
		return nil, fmt.Errorf("not found")
	}
	fmt.Printf("username: %s, host: %s\n", username, host)

	resp, err := http.Get(fmt.Sprintf("http://admin:super-secret@couchdb-service.default.svc.cluster.local:5984/users/%s", username))
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not found")
	}

	return &Webfinger{
		Subject: fmt.Sprintf("acct:%s@%s", username, host),
		Aliases: []string{
			fmt.Sprintf("http://%s/u/%s", host, username),
		},
	}, nil
}
