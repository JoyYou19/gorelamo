package database

import (
	"fmt"
	"gorelamo/transport"
)

type Database struct {
	http        *transport.HTTP
	baseURL     string
	name        string
	documentTag string
}

func New(http *transport.HTTP, baseURL, name string) *Database {
	return &Database{
		http:        http,
		baseURL:     baseURL,
		name:        name,
		documentTag: "document",
	}
}

func (db *Database) endpoint(path string) string {
	return fmt.Sprintf("%s/databases/%s/%s", db.baseURL, db.name, path)
}

func (db *Database) WithDocumentTag(tag string) *Database {
	if tag != "" {
		db.documentTag = tag
	}
	return db
}
