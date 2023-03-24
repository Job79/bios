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

// GetShowings returns a list with showings
// Query params: uid, movie, room, limit, offset, load
func (ctx Context) GetShowings(w http.ResponseWriter, r *http.Request) {
	// Parse query params
	opts := store.ShowingOptions{Limit: 25}
	form.Unmarshal(r.URL.Query(), &opts)

	// Fetch showings
	results, err := ctx.DB.FetchShowings(opts)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(results) == 0 {
		Error(w, http.StatusNotFound, errors.New("no showings found"))
		return
	}

	// Hydrate room if requested
	if slices.Contains(r.URL.Query()["load"], "room") {
		results, err = ctx.DB.HydrateShowingRoom(results)
		if err != nil {
			Error(w, http.StatusInternalServerError, err)
			return
		}
	}

	// Hydrate movie if requested
	if slices.Contains(r.URL.Query()["load"], "movie") {
		results, err = ctx.DB.HydrateShowingMovie(results)
		if err != nil {
			Error(w, http.StatusInternalServerError, err)
			return
		}
	}

	Json(w, http.StatusOK, results)
}

// TODO: add editShowing

// DeleteShowings removes showings
// Body: []model.Showing
func (ctx Context) DeleteShowings(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var showings []*model.Showing
	if err := json.NewDecoder(r.Body).Decode(&showings); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	// Delete showings
	if err := ctx.DB.DeleteShowings(showings); err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, map[string]bool{"acknowledged": true})
}
