package store

import (
	"bios/store/model"
	"fmt"
)

// ClassificationOptions contain the options for fetching classifications
type ClassificationOptions struct {
	ID  []int
	UID []string `form:"uid"`

	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

// FetchClassifications fetches classification from the database based on given options
func (s Store) FetchClassifications(o ClassificationOptions) (results []*model.Movie, err error) {
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
		"SELECT id, uid, name FROM classification %s LIMIT ? OFFSET ?",
		whereQuery(conditions),
	)

	// Execute query and collect results
	return results, s.db.Select(
		&results,
		s.db.Rebind(query),
		append(args, o.Limit, o.Offset)...,
	)
}

// AttachClassifications detects the ids of given objects based on the uid's
// Should be called before flushing when doing updates
func (s Store) AttachClassifications(classifications []*model.Classification) ([]*model.Classification, error) {
	return attachGeneric(s, "classification", classifications)
}

// FlushClassifications flushes given classifications to the database
func (s Store) FlushClassifications(classifications []*model.Classification, insert, update bool) (results []*model.Classification, err error) {
	if insert {
		// Collect all classifications without id, these are new
		inserts := filter(classifications, func(c *model.Classification) bool { return c.ID == 0 })
		if len(inserts) > 0 {
			// Build query, execute and collect results
			query, queryArgs, _ := s.db.BindNamed("INSERT INTO classification (name, description) VALUES (:name, :description) RETURNING id, uid, name, description", inserts)
			var insertResults []*model.Classification
			if err := s.db.Select(&insertResults, s.db.Rebind(query), queryArgs...); err != nil {
				return nil, err
			}
			results = append(results, insertResults...)
		}
	}

	if update {
		// Collect all classifications with id, these already exist and should be updated
		updates := filter(classifications, func(c *model.Classification) bool { return c.ID != 0 })
		if len(updates) > 0 {
			// Build query, execute and collect results
			query, queryArgs, _ := s.db.BindNamed("INSERT INTO classification (id, uid, name, description) VALUES (:id,:uid,:name,:description)", updates)
			query = s.db.Rebind(query + " ON CONFLICT (id) DO UPDATE SET name = excluded.name, description = excluded.description RETURNING id, uid, name, description")
			var updateResults []*model.Classification
			if err := s.db.Select(&updateResults, query, queryArgs...); err != nil {
				return nil, err
			}
			results = append(results, updateResults...)
		}
	}

	return results, nil
}

// DeleteClassifications removes given classifications from the database
func (s Store) DeleteClassifications(classifications []*model.Classification) error {
	return deleteGeneric(s, "classification", classifications)
}
