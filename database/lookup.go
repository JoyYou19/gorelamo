package database

import "gorelamo/types"

func (db *Database) Lookup(id any) (types.Document, error) {
	payload := map[string]any{
		"id": id,
	}

	var doc types.Document
	err := db.http.Post(
		db.endpoint("lookup"),
		payload,
		&doc,
	)

	return doc, err
}
