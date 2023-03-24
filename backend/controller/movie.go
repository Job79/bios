package controller

import (
	"bios/service/form"
	"bios/store"
	"bios/store/model"
	"encoding/json"
	"errors"
	"golang.org/x/exp/slices"
	"net/http"
)

// GetMovies returns a list with movies
// Query params: uid, genre, status, query, limit, offset, load
func (ctx Context) GetMovies(w http.ResponseWriter, r *http.Request) {
	// Parse query params
	opts := store.MovieOptions{Limit: 25}
	form.Unmarshal(r.URL.Query(), &opts)

	// Fetch movies
	results, err := ctx.DB.FetchMovies(opts)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(results) == 0 {
		Error(w, http.StatusNotFound, errors.New("no movies found"))
		return
	}

	// Hydrate genres if requested
	if slices.Contains(r.URL.Query()["load"], "genres") {
		results, err = ctx.DB.HydrateMovieGenres(results)
		if err != nil {
			Error(w, http.StatusInternalServerError, err)
			return
		}
	}

	// Hydrate classifications if requested
	if slices.Contains(r.URL.Query()["load"], "classifications") {
		results, err = ctx.DB.HydrateMovieClassifications(results)
		if err != nil {
			Error(w, http.StatusInternalServerError, err)
			return
		}
	}

	// Hydrate files if requested
	if slices.Contains(r.URL.Query()["load"], "files") {
		results, err = ctx.DB.HydrateMovieFiles(results)
		if err != nil {
			Error(w, http.StatusInternalServerError, err)
			return
		}
	}

	Json(w, http.StatusOK, results)
}

// EditMovies creates or updates movies
// Body: []model.Movie
func (ctx Context) EditMovies(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var movies []*model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movies); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	// Verify movies
	for _, m := range movies {
		if err := m.Verify(); err != nil {
			Error(w, http.StatusBadRequest, err)
			return
		}
	}

	// Detect if movies are new or already exist
	movies, err := ctx.DB.AttachMovies(movies)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	// Save or update movies
	insert := r.Method == "POST"
	update := r.Method == "POST" || r.Method == "PUT"
	movies, err = ctx.DB.FlushMovies(movies, insert, update)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	// TODO: save movie relationships

	Json(w, http.StatusOK, movies)
}

// DeleteMovies removes movies
// Body: []model.Movie
func (ctx Context) DeleteMovies(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var movies []*model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movies); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	// Delete movies
	if err := ctx.DB.DeleteMovies(movies); err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, map[string]bool{"acknowledged": true})
}
