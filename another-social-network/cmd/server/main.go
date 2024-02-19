// cmd/server/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/martcl/another-social-network/pkg/db"

	"github.com/martcl/another-social-network/api/router"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	ClientKey = contextKey("client")
)

func main() {
	port := 3000

	client, err := db.NewClient()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := db.NewDBContext(context.Background(), client)

	r := router.NewRouter(ctx)

	fmt.Printf("Server running on port %d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
