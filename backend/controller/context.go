package controller

import (
	"bios/config"
	"bios/store"
	"encoding/json"
	"net/http"
)

// Context is used to pass around a global context
// All controllers should be declared against this type
type Context struct {
	DB   store.Store
	Conf config.Config
}

// Error returns an error message to the client
func Error(w http.ResponseWriter, status int, err error) {
	Json(w, status, map[string]string{"error": err.Error()})
}

// Json returns a json response to the client
func Json(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
