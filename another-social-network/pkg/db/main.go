package db

import (
	"context"

	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb" // The CouchDB driver
)

// NewClient creates a new Kivik client
func NewClient() (*kivik.Client, error) {
	return kivik.New("couch", "http://admin:super-secret@couchdb-service.default.svc.cluster.local:5984")
}

// User represents a user in the database
type User struct {
	ID       string `json:"_id"`
	Rev      string `json:"_rev,omitempty"`
	Username string `json:"username"`
}

// Save saves the user to the database
func (u *User) Save(client *kivik.Client) error {
	db := client.DB("users")

	rev, err := db.Put(context.Background(), u.ID, u)

	if err != nil {
		return err
	}
	u.Rev = rev
	return nil
}

type DBContext struct {
	context.Context
	Client *kivik.Client
}

func NewDBContext(ctx context.Context, client *kivik.Client) *DBContext {
	return &DBContext{
		Context: ctx,
		Client:  client,
	}
}
