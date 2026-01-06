package database

func (db *Database) Insert(data map[string]any) error {
	payload := map[string]any{
		db.documentTag: data,
	}

	return db.http.Post(
		db.endpoint("insert"),
		payload,
		nil,
	)
}

func (db *Database) InsertMany(docs []map[string]any) error {
	payload := make([]map[string]any, 0, len(docs))

	for _, d := range docs {
		payload = append(payload, map[string]any{
			db.documentTag: d,
		})
	}

	return db.http.Post(
		db.endpoint("insert"),
		payload,
		nil,
	)
}
