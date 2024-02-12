package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type Webfinger struct {
	Subject string   `json:"subject"`
	Aliases []string `json:"aliases"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /.well-known/webfinger", GETWebfinger)
	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe(":3000", mux)
}

func GETWebfinger(w http.ResponseWriter, r *http.Request) {
	resource := r.URL.Query().Get("resource")

	// check the resource starts with acct:
	if len(resource) < 25 || resource[:5] != "acct:" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resource = resource[5:] // remove "acct:"
	usernameHost := strings.Split(resource, "@")
	if len(usernameHost) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := usernameHost[1]
	host := usernameHost[2]

	// check if the host is the host of the server
	if host != "social-network.local" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if the user exists
	resp, err := http.Get(fmt.Sprintf("http://admin:super-secret@couchdb.local/users/%s", username))
	if err != nil || resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	webfinger := Webfinger{
		Subject: resource,
		Aliases: []string{
			"http://social-network.local/u/martin",
		},
	}

	jsonResponse, err := json.Marshal(webfinger)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
