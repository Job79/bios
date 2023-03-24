package store

import (
	"bios/store/model"
	"fmt"
)

// GenreOptions contain the options for fetching genres
type GenreOptions struct {
	ID  []int
	UID []string `form:"uid"`

	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

// FetchGenres fetches genres from the database based on given options
func (s Store) FetchGenres(o GenreOptions) (results []model.Genre, err error) {
	var (
		conditions = make([]string, 0)
		args       = make([]any, 0)
	)

	// Collect arguments based on options
	if len(o.ID) > 0 {
		conditions = append(conditions, "id IN ("+inQuery(len(o.ID))+")")
		args = append(args, mapSliceToAny(o.ID)...)
	} else if len(o.UID) > 0 {
		conditions = append(conditions, "uid IN ("+inQuery(len(o.UID))+")")
		args = append(args, mapSliceToAny(o.UID)...)
	}

	// Build query
	query := fmt.Sprintf(
		"SELECT id, uid, name FROM genre %s LIMIT ? OFFSET ?",
		whereQuery(conditions),
	)

	// Execute query and collect results
	return results, s.db.Select(
		&results,
		s.db.Rebind(query),
		append(args, o.Limit, o.Offset)...,
	)
}

// AttachGenres detects the ids of given objects based on the uid's
// Should be called before flushing when doing updates
func (s Store) AttachGenres(genres []*model.Genre) ([]*model.Genre, error) {
	return attachGeneric(s, "genre", genres)
}

// FlushGenres flushes given genres to the database
func (s Store) FlushGenres(genres []*model.Genre, insert, update bool) (results []*model.Genre, err error) {
	if insert {
		// Collect all genres without id, these are new
		inserts := filter(genres, func(g *model.Genre) bool { return g.ID == 0 })
		if len(inserts) > 0 {
			// Build query, execute and collect results
			query, queryArgs, _ := s.db.BindNamed("INSERT INTO genre (name) VALUES (:name) RETURNING id, uid, name", inserts)
			var insertResults []*model.Genre
			if err := s.db.Select(&insertResults, s.db.Rebind(query), queryArgs...); err != nil {
				return nil, err
			}
			results = append(results, insertResults...)
		}
	}

	if update {
		// Collect all genres with id, these already exist and should be updated
		updates := filter(genres, func(g *model.Genre) bool { return g.ID != 0 })
		if len(updates) > 0 {
			// Build query, execute and collect results
			query, queryArgs, _ := s.db.BindNamed("INSERT INTO genre (id, uid, name) VALUES (:id,:uid,:name)", updates)
			query = s.db.Rebind(query + " ON CONFLICT (id) DO UPDATE SET name = excluded.name RETURNING id, uid, name")
			var updateResults []*model.Genre
			if err := s.db.Select(&updateResults, query, queryArgs...); err != nil {
				return nil, err
			}
			results = append(results, updateResults...)
		}
	}

	return results, nil
}

// DeleteGenres removes given genres from the database
func (s Store) DeleteGenres(genres []*model.Genre) error {
	return deleteGeneric(s, "genre", genres)
}
