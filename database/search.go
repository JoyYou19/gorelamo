package database

import (
	"fmt"
	"github.com/JoyYou19/gorelamo/types"
)

type SearchQuery struct {
	db     *Database
	query  string
	docs   int
	offset int
}

func (db *Database) Search(q string) *SearchQuery {
	return &SearchQuery{
		db:    db,
		query: q,
	}
}

func (s *SearchQuery) Limit(n int) *SearchQuery {
	s.docs = n
	return s
}

func (s *SearchQuery) Offset(n int) *SearchQuery {
	s.offset = n
	return s
}

func (s *SearchQuery) Do() (*types.SearchResult, error) {
	xmlPayload := fmt.Appendf(nil,
		"<query>%s</query>",
		s.query,
	)

	var result types.SearchResult
	err := s.db.http.PostXML(
		s.db.endpoint("search"),
		xmlPayload,
		&result,
	)

	return &result, err
}
