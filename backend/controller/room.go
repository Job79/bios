package controller

import (
	"bios/service/form"
	"bios/store"
	"bios/store/model"
	"encoding/json"
	"errors"
	"net/http"
)

// GetRooms returns a list with rooms
// Query params: uid, limit, offset
func (ctx Context) GetRooms(w http.ResponseWriter, r *http.Request) {
	// Parse query params
	opts := store.RoomOptions{Limit: 25}
	form.Unmarshal(r.URL.Query(), &opts)

	// Fetch rooms
	results, err := ctx.DB.FetchRooms(opts)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(results) == 0 {
		Error(w, http.StatusNotFound, errors.New("no rooms found"))
		return
	}

	Json(w, http.StatusOK, results)
}

// EditRooms creates or updates rooms
// Body: []model.Room
func (ctx Context) EditRooms(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var rooms []*model.Room
	if err := json.NewDecoder(r.Body).Decode(&rooms); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	// Verify rooms
	for _, r := range rooms {
		if err := r.Verify(); err != nil {
			Error(w, http.StatusBadRequest, err)
			return
		}
	}

	// Detect if rooms are new or already exist
	rooms, err := ctx.DB.AttachRooms(rooms)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	// Save or update rooms
	insert := r.Method == "POST"
	update := r.Method == "POST" || r.Method == "PUT"
	rooms, err = ctx.DB.FlushRooms(rooms, insert, update)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, rooms)
}

// DeleteRooms removes rooms
// Body: []model.Room
func (ctx Context) DeleteRooms(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var rooms []*model.Room
	if err := json.NewDecoder(r.Body).Decode(&rooms); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	// Delete rooms
	if err := ctx.DB.DeleteRooms(rooms); err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, map[string]bool{"acknowledged": true})
}
