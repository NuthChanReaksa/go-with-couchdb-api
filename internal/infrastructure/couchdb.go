package infrastructure

import (
	"context"

	"github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
)

func NewCouchDBConnection(ctx context.Context, dsn string) (*kivik.Client, error) {
	client, err := kivik.New("couch", dsn)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func OpenDatabase(client *kivik.Client, dbName string) (*kivik.DB, error) {
	db := client.DB(dbName, nil) // Remove `ctx` and pass `nil` for options if not using any
	if err := db.Err(); err != nil {
		return nil, err
	}
	return db, nil
}
