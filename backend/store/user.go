package store

import (
	"bios/store/model"
)

// UserOptions contain the options for fetching users
type UserOptions struct {
	ID   []int
	UID  []string
	Name []string
}

// FetchUsers fetches users from the database based on given options
func (s Store) FetchUsers(o UserOptions) (results []model.User, err error) {
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
	} else if len(o.Name) > 0 {
		conditions = append(conditions, "name IN ("+inQuery(len(o.Name))+")")
		args = append(args, mapSliceToAny(o.Name)...)
	}

	// Build query, execute and collect results
	return results, s.db.Select(
		&results,
		s.db.Rebind("SELECT id, uid, name, password FROM \"user\" "+whereQuery(conditions)),
		args...,
	)
}
