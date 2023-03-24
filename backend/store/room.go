package store

import (
	"bios/store/model"
	"fmt"
)

// RoomOptions contain the options for fetching rooms
type RoomOptions struct {
	ID  []int
	UID []string `form:"uid"`

	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

// FetchRooms fetches rooms from the database based on given options
func (s Store) FetchRooms(o RoomOptions) (results []model.Room, err error) {
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
		"SELECT id, uid, code FROM room %s LIMIT ? OFFSET ?",
		whereQuery(conditions),
	)

	// Execute query and collect results
	return results, s.db.Select(
		&results,
		s.db.Rebind(query),
		append(args, o.Limit, o.Offset)...,
	)
}

// AttachRooms detects the ids of given objects based on the uid's
// Should be called before flushing when doing updates
func (s Store) AttachRooms(rooms []*model.Room) ([]*model.Room, error) {
	return attachGeneric(s, "room", rooms)
}

// FlushRooms flushes given genres to the database
func (s Store) FlushRooms(rooms []*model.Room, insert, update bool) (results []*model.Room, err error) {
	if insert {
		// Collect all rooms without id, these are new
		inserts := filter(rooms, func(r *model.Room) bool { return r.ID == 0 })
		if len(inserts) > 0 {
			// Build query, execute and collect results
			query, queryArgs, _ := s.db.BindNamed("INSERT INTO room (code) VALUES (:code) RETURNING id, uid, code", inserts)
			var insertResults []*model.Room
			if err := s.db.Select(&insertResults, s.db.Rebind(query), queryArgs...); err != nil {
				return nil, err
			}
			results = append(results, insertResults...)
		}
	}

	if update {
		// Collect all rooms with id, these already exist and should be updated
		updates := filter(rooms, func(r *model.Room) bool { return r.ID != 0 })
		if len(updates) > 0 {
			// Build query, execute and collect results
			query, queryArgs, _ := s.db.BindNamed("INSERT INTO room (id, uid, code) VALUES (:id,:uid,:code)", updates)
			query = s.db.Rebind(query + " ON CONFLICT (id) DO UPDATE SET code = excluded.code RETURNING id, uid, code")
			var updateResults []*model.Room
			if err := s.db.Select(&updateResults, query, queryArgs...); err != nil {
				return nil, err
			}
			results = append(results, updateResults...)
		}
	}

	return results, nil
}

// DeleteRooms removes given rooms from the database
func (s Store) DeleteRooms(rooms []*model.Room) error {
	return deleteGeneric(s, "room", rooms)
}
