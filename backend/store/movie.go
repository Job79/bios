package store

import (
	"bios/store/model"
	"fmt"
	"strings"
)

// MovieOptions contain the options for fetching movies
type MovieOptions struct {
	ID     []int
	UID    []string `form:"uid"`
	Genre  []string `form:"genre"`
	Status []string `form:"status"`
	Query  string   `form:"query"`

	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

// FetchMovies fetches movies from the database based on given options
func (s Store) FetchMovies(o MovieOptions) (results []*model.Movie, err error) {
	var (
		conditions = make([]string, 0)
		joins      = make([]string, 0)
		orderBy    = make([]string, 0)
		args       = make([]any, 0)
	)

	// Collect arguments based on options
	if len(o.ID) > 0 {
		conditions = append(conditions, "m.id IN ("+inQuery(len(o.ID))+")")
		args = append(args, mapSliceToAny(o.ID)...)
	} else if len(o.UID) > 0 {
		conditions = append(conditions, "m.uid IN ("+inQuery(len(o.UID))+")")
		args = append(args, mapSliceToAny(o.UID)...)
	} else if len(o.Status) > 0 {
		conditions = append(conditions, "m.status IN ("+inQuery(len(o.Status))+")")
		args = append(args, mapSliceToAny(o.Status)...)
	} else if len(o.Genre) > 0 {
		joins = append(
			joins,
			"INNER JOIN movie_genre mg ON m.id = mg.movie_id",
			"INNER JOIN genre g ON g.id = mg.genre_id",
		)
		conditions = append(conditions, "g.name IN ("+inQuery(len(o.Genre))+")")
		args = append(args, mapSliceToAny(o.Genre)...)
	} else if o.Query != "" {
		conditions = append(conditions, "to_tsvector(m.name) @@ plainto_tsquery(?)")
		orderBy = append(orderBy, "ts_rank_cd(to_tsvector(m.name), plainto_tsquery(?)) DESC")
		args = append(args, o.Query, o.Query)
	}

	// Build query
	query := fmt.Sprintf(
		`SELECT m.id, m.uid, m.name, m.description, m.total_time, m.release_date, m.status
					FROM movie m %s %s
					GROUP BY m.id %s
					LIMIT ? OFFSET ?`,
		strings.Join(joins, "\n"),
		whereQuery(conditions),
		orderByQuery(orderBy),
	)

	// Execute query and collect results
	return results, s.db.Select(
		&results,
		s.db.Rebind(query),
		append(args, o.Limit, o.Offset)...,
	)
}

// HydrateMovieGenres fetches the genres for all given movies
func (s Store) HydrateMovieGenres(movies []*model.Movie) ([]*model.Movie, error) {
	// Collect movie IDs and fetch genres
	query := "SELECT g.id, g.uid, g.name, mg.movie_id FROM genre g INNER JOIN movie_genre mg ON g.id = mg.genre_id WHERE mg.movie_id IN (" + inQuery(len(movies)) + ")"
	row, err := s.db.Queryx(
		s.db.Rebind(query),
		mapSlice(movies, func(m *model.Movie) any { return m.ID })...,
	)
	if err != nil {
		return movies, err
	}

	// Create map[movieID]movieIdx
	idxMap := mapToIdx(movies, func(m *model.Movie) int { return m.ID })
	for row.Next() {
		// Fetch genre
		genre, movieID := &model.Genre{}, 0
		if err := row.Scan(&genre.ID, &genre.UID, &genre.Name, &movieID); err != nil {
			return movies, err
		}

		// Add genre to movie
		movies[idxMap[movieID]].Genres = append(movies[idxMap[movieID]].Genres, genre)
	}
	return movies, row.Err()
}

// HydrateMovieClassifications fetches the classifications for all given movies
func (s Store) HydrateMovieClassifications(movies []*model.Movie) ([]*model.Movie, error) {
	// Collect movie IDs and fetch classifications
	query := "SELECT c.id, c.uid, c.name, c.description, mc.movie_id FROM classification c INNER JOIN movie_classification mc ON c.id = mc.classification_id WHERE mc.movie_id IN (" + inQuery(len(movies)) + ")"
	row, err := s.db.Queryx(
		s.db.Rebind(query),
		mapSlice(movies, func(m *model.Movie) any { return m.ID })...,
	)
	if err != nil {
		return movies, err
	}

	// Create map[movieID]movieIdx
	idxMap := mapToIdx(movies, func(m *model.Movie) int { return m.ID })
	for row.Next() {
		// Fetch classification
		classification, movieID := &model.Movie{}, 0
		if err := row.Scan(&classification.ID, &classification.UID, &classification.Name, &classification.Description, &movieID); err != nil {
			return movies, err
		}

		// Add classification to movie
		movies[idxMap[movieID]].Classifications = append(movies[idxMap[movieID]].Classifications, classification)
	}
	return movies, row.Err()
}

// HydrateMovieFiles fetches the files for all given movies
func (s Store) HydrateMovieFiles(movies []*model.Movie) ([]*model.Movie, error) {
	// Collect movie IDs and fetch files
	query := "SELECT f.id, f.uid, f.path, f.type, mf.movie_id FROM file f INNER JOIN movie_file mf ON f.id = mf.file_id WHERE mf.movie_id IN (" + inQuery(len(movies)) + ") ORDER BY mf.position"
	row, err := s.db.Queryx(
		s.db.Rebind(query),
		mapSlice(movies, func(m *model.Movie) any { return m.ID })...,
	)
	if err != nil {
		return movies, err
	}

	// Create map[movieID]movieIdx
	idxMap := mapToIdx(movies, func(m *model.Movie) int { return m.ID })
	for row.Next() {
		// Fetch file
		file, movieID := &model.File{}, 0
		if err := row.Scan(&file.ID, &file.UID, &file.Path, &file.Type, &movieID); err != nil {
			return movies, err
		}

		// Add file to movie
		movies[idxMap[movieID]].Files = append(movies[idxMap[movieID]].Files, file)
	}
	return movies, row.Err()
}

// AttachMovies detects the ids of given objects based on the uid's
// Should be called before flushing when doing updates
func (s Store) AttachMovies(movies []*model.Movie) ([]*model.Movie, error) {
	return attachGeneric(s, "movie", movies)
}

// FlushMovies flushes given movies to the database
func (s Store) FlushMovies(movies []*model.Movie, insert, update bool) (results []*model.Movie, err error) {
	if insert {
		// Collect all movies without id, these are new
		inserts := filter(movies, func(m *model.Movie) bool { return m.ID == 0 })
		if len(inserts) > 0 {
			// Build query, execute and collect results
			query, queryArgs, _ := s.db.BindNamed(
				`INSERT INTO movie (name, description, total_time, release_date, status) 
						VALUES (:name, :description, :total_time, :release_date, :status)
						RETURNING id, uid, name, description, total_time, release_date, status`,
				inserts,
			)
			var insertResults []*model.Movie
			if err := s.db.Select(&insertResults, s.db.Rebind(query), queryArgs...); err != nil {
				return nil, err
			}
			results = append(results, insertResults...)
		}
	}

	if update {
		// Collect all movies with id, these already exist and should be updated
		updates := filter(movies, func(m *model.Movie) bool { return m.ID != 0 })
		if len(updates) > 0 {
			// Build query, execute and collect results
			query, queryArgs, _ := s.db.BindNamed(
				`INSERT INTO movie (id, name, description, total_time, release_date, status)
						VALUES (:id, :name, :description, :total_time, :release_date, :status)`,
				updates,
			)
			query = s.db.Rebind(
				query + ` ON CONFLICT (id)
				DO UPDATE SET name = excluded.name, description = excluded.description, total_time = excluded.total_time, release_date = excluded.release_date, status = excluded.status
				RETURNING id, uid, name, description, total_time, release_date, status`,
			)
			var updateResults []*model.Movie
			if err := s.db.Select(&updateResults, query, queryArgs...); err != nil {
				return nil, err
			}
			results = append(results, updateResults...)
		}
	}

	return results, nil
}

// FlushMovieGenres flushes the relation to the genres for all given movies
func (s Store) FlushMovieGenres(movies []*model.Movie) error {
	// Fetch the relations that are currently stored in the database
	storedRelations, err := fetchRelationsMap(s, "movie_genre", "movie_id", "genre_id", mapSlice(movies, func(m *model.Movie) int { return m.ID }))
	if err != nil {
		return err
	}

	// Compare the stored relations with the new relations
	removeArgs, insertArgs := make([]any, 0), make([]any, 0)
	for _, movie := range movies {
		relations := mapSlice(movie.Genres, func(g *model.Genre) int { return g.ID })

		// Remove all relations that are not in the new relations
		for _, genreID := range difference(storedRelations[movie.ID], relations) {
			removeArgs = append(removeArgs, movie.ID, genreID)
		}

		// Insert all relations that are not in the stored relations
		for _, genreID := range difference(relations, storedRelations[movie.ID]) {
			insertArgs = append(insertArgs, movie.ID, genreID)
		}
	}

	// Build and execute queries
	if len(removeArgs) > 0 {
		if _, err = s.db.Query(
			"DELETE FROM movie_genre WHERE "+strings.Repeat("movie_id = ? AND genre_id = ? OR ", len(removeArgs)/2-1)+"movie_id = ? AND genre_id = ?",
			removeArgs...,
		); err != nil {
			return err
		}
	}
	if len(insertArgs) > 0 {
		query, queryArgs, _ := s.db.BindNamed("INSERT INTO movie_genre (movie_id, genre_id) VALUES (:movie_id, :genre_id)", insertArgs)
		if _, err = s.db.Query(s.db.Rebind(query), queryArgs...); err != nil {
			return err
		}
	}
	return nil
}

// DeleteMovies removes given movies from the database
func (s Store) DeleteMovies(movies []*model.Movie) error {
	return deleteGeneric(s, "movie", movies)
}
