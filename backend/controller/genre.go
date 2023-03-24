package controller

import (
	"bios/service/form"
	"bios/store"
	"bios/store/model"
	"encoding/json"
	"errors"
	"net/http"
)

// GetGenres returns a list with classifications
// Query params: uid, limit, offset
func (ctx Context) GetGenres(w http.ResponseWriter, r *http.Request) {
	// Parse query params
	opts := store.GenreOptions{Limit: 25}
	form.Unmarshal(r.URL.Query(), &opts)

	// Fetch genres
	results, err := ctx.DB.FetchGenres(opts)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(results) == 0 {
		Error(w, http.StatusNotFound, errors.New("no genres found"))
		return
	}

	Json(w, http.StatusOK, results)
}

// EditGenres creates or updates genres
// Body: []model.Genre
func (ctx Context) EditGenres(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var genres []*model.Genre
	if err := json.NewDecoder(r.Body).Decode(&genres); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	// Verify genres
	for _, g := range genres {
		if err := g.Verify(); err != nil {
			Error(w, http.StatusBadRequest, err)
			return
		}
	}

	// Detect if genres are new or already exist
	genres, err := ctx.DB.AttachGenres(genres)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	// Save or update genres
	insert := r.Method == "POST"
	update := r.Method == "POST" || r.Method == "PUT"
	genres, err = ctx.DB.FlushGenres(genres, insert, update)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, genres)
}

// DeleteGenres removes genres
// Body: []model.Genre
func (ctx Context) DeleteGenres(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var genres []*model.Genre
	if err := json.NewDecoder(r.Body).Decode(&genres); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	// Delete genres
	if err := ctx.DB.DeleteGenres(genres); err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, map[string]bool{"acknowledged": true})
}
