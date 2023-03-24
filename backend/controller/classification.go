package controller

import (
	"bios/service/form"
	"bios/store"
	"bios/store/model"
	"encoding/json"
	"errors"
	"net/http"
)

// GetClassifications returns a list with classifications
// Query params: uid, limit, offset
func (ctx Context) GetClassifications(w http.ResponseWriter, r *http.Request) {
	// Parse query params
	opts := store.ClassificationOptions{Limit: 25}
	form.Unmarshal(r.URL.Query(), &opts)

	// Fetch classifications
	results, err := ctx.DB.FetchClassifications(opts)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(results) == 0 {
		Error(w, http.StatusNotFound, errors.New("no classifications found"))
		return
	}

	Json(w, http.StatusOK, results)
}

// EditClassifications creates or updates classifications
// Body: []model.Classification
func (ctx Context) EditClassifications(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var classifications []*model.Classification
	if err := json.NewDecoder(r.Body).Decode(&classifications); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	// Verify classifications
	for _, c := range classifications {
		if err := c.Verify(); err != nil {
			Error(w, http.StatusBadRequest, err)
			return
		}
	}

	// Detect if classifications are new or already exist
	classifications, err := ctx.DB.AttachClassifications(classifications)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	// Save or update classifications
	insert := r.Method == "POST"
	update := r.Method == "POST" || r.Method == "PUT"
	classifications, err = ctx.DB.FlushClassifications(classifications, insert, update)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, classifications)
}

// DeleteClassifications removes classifications
// Body: []model.Classification
func (ctx Context) DeleteClassifications(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var classifications []*model.Classification
	if err := json.NewDecoder(r.Body).Decode(&classifications); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	// Delete classifications
	if err := ctx.DB.DeleteClassifications(classifications); err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, map[string]bool{"acknowledged": true})
}
