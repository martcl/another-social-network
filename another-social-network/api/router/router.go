// api/router/router.go
package router

import (
	"net/http"

	"github.com/martcl/another-social-network/pkg/db"
)

func NewRouter(ctx *db.DBContext) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello, world!"))
	})

	r.HandleFunc("GET /.well-known/webfinger", func(w http.ResponseWriter, r *http.Request) {
		resource := r.URL.Query().Get("resource")
		if resource == "" {
			http.Error(w, "resource query parameter is required", http.StatusBadRequest)
			return
		}

		var user *db.User
		err := ctx.Client.DB("users").Get(ctx, resource).ScanDoc(&user)

		if err != nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		r.Header.Set("Content-Type", "application/json")
		w.Write([]byte(`{"subject": "acct:` + user.Username + `@localhost"}`))
	})

	return r
}
