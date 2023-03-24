package store

import (
	"bios/store/model"
	"fmt"
	"strings"
)

// ShowingOptions contain the options for fetching showings
type ShowingOptions struct {
	ID       []int
	UID      []string `form:"uid"`
	MovieUID []string `form:"movie"`
	RoomUID  []string `form:"room"`

	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

// FetchShowings fetches showings from the database based on given options
func (s Store) FetchShowings(o ShowingOptions) (results []model.Showing, err error) {
	var (
		conditions = make([]string, 0)
		joins      = make([]string, 0)
		args       = make([]any, 0)
	)

	// Collect arguments based on options
	if len(o.ID) > 0 {
		conditions = append(conditions, "s.id IN ("+inQuery(len(o.ID))+")")
		args = append(args, mapSliceToAny(o.ID)...)
	} else if len(o.UID) > 0 {
		conditions = append(conditions, "s.uid IN ("+inQuery(len(o.UID))+")")
		args = append(args, mapSliceToAny(o.UID)...)
	} else if len(o.MovieUID) > 0 {
		joins = append(joins, "INNER JOIN movie m ON m.id = s.movie_id")
		conditions = append(conditions, "m.uid IN ("+inQuery(len(o.MovieUID))+")")
		args = append(args, mapSliceToAny(o.MovieUID)...)
	} else if len(o.RoomUID) > 0 {
		joins = append(joins, "INNER JOIN room r ON r.id = s.room_id")
		conditions = append(conditions, "r.uid IN ("+inQuery(len(o.RoomUID))+")")
		args = append(args, mapSliceToAny(o.RoomUID)...)
	}

	// Build query
	query := fmt.Sprintf(
		`SELECT s.id, s.uid, s.movie_id, s.room_id, s.start_time, s.end_time
					FROM showing s %s %s
					ORDER BY s.start_time
					LIMIT ? OFFSET ?`,
		strings.Join(joins, "\n"),
		whereQuery(conditions),
	)

	// Execute query and collect results
	return results, s.db.Select(
		&results,
		s.db.Rebind(query),
		append(args, o.Limit, o.Offset)...,
	)
}

// HydrateShowingMovie fetches the movie for all given showings
func (s Store) HydrateShowingMovie(showings []model.Showing) ([]model.Showing, error) {
	// Collect movie IDs and fetch movies
	opts := MovieOptions{ID: mapSlice(showings, func(s model.Showing) int { return s.MovieID }), Limit: len(showings)}
	movies, err := s.FetchMovies(opts)
	if err != nil {
		return showings, err
	}

	// Create map[movieID]movieIdx
	idxMap := mapToIdx(movies, func(m *model.Movie) int { return m.ID })
	for idx := range showings {
		idx, ok := idxMap[showings[idx].MovieID] // Get movieIdx by ID from idxMap
		if ok {
			showings[idx].Movie = movies[idx]
		}
	}

	return showings, nil
}

// HydrateShowingRoom fetches the room for all given showings
func (s Store) HydrateShowingRoom(showings []model.Showing) ([]model.Showing, error) {
	// Collect room IDs and fetch rooms
	opts := RoomOptions{ID: mapSlice(showings, func(s model.Showing) int { return s.RoomID }), Limit: len(showings)}
	rooms, err := s.FetchRooms(opts)
	if err != nil {
		return showings, err
	}

	// Create map[roomID]roomIdx
	idxMap := mapToIdx(rooms, func(m model.Room) int { return m.ID })
	for idx := range showings {
		idx, ok := idxMap[showings[idx].RoomID] // Get showingIdx by ID from idxMap
		if ok {
			showings[idx].Room = &rooms[idx]
		}
	}

	return showings, nil
}

// AttachShowings detects the ids of given objects based on the uid's
// Should be called before flushing when doing updates
func (s Store) AttachShowings(showings []*model.Showing) ([]*model.Showing, error) {
	return attachGeneric(s, "showing", showings)
}

func (s Store) FlushShowings(showings []*model.Showing, insert, update bool) (results []*model.Showing, err error) {
	if insert {
		// Collect all showings without id, these are new
		inserts := filter(showings, func(s *model.Showing) bool { return s.ID == 0 })
		if len(inserts) > 0 { // Abort when non are found
			// Build query, execute and collect results
			query, queryArgs, _ := s.db.BindNamed("INSERT INTO showing (movie_id, room_id, start_time, end_time, uid) VALUES (:movie_id, :room_id, :start_time, :end_time) RETURNING id, uid, code", inserts)
			var insertResults []*model.Showing
			if err := s.db.Select(&insertResults, s.db.Rebind(query), queryArgs...); err != nil {
				return nil, err
			}
			results = append(results, insertResults...)
		}
	}

	if update {
		// Collect all showings with id, these already exist and should be updated
		updates := filter(showings, func(s *model.Showing) bool { return s.ID != 0 })
		if len(updates) > 0 {
			// Build query, execute and collect results
			query, queryArgs, _ := s.db.BindNamed("INSERT INTO room (id, uid, code) VALUES (:id,:uid,:code)", updates)
			query = s.db.Rebind(query + " ON CONFLICT (id) DO UPDATE SET code = excluded.code RETURNING id, uid, code")
			var updateResults []*model.Showing
			if err := s.db.Select(&updateResults, query, queryArgs...); err != nil {
				return nil, err
			}
			results = append(results, updateResults...)
		}
	}

	return results, nil
}

// DeleteShowings removes given showings from the database
func (s Store) DeleteShowings(showings []*model.Showing) error {
	return deleteGeneric(s, "showing", showings)
}
