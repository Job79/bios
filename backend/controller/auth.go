package controller

import (
	"bios/store"
	"encoding/base64"
	"errors"
	"net/http"
)

// PostLogin lets the user authenticate and returns a token
// Form params: name, password
func (ctx Context) PostLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		Error(w, http.StatusBadRequest, err)
	}

	//  Fetch the user
	opts := store.UserOptions{}
	opts.Name = []string{r.PostForm.Get("user")}

	results, err := ctx.DB.FetchUsers(opts)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(results) != 1 {
		Error(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	// Verify the password
	if ok, _ := results[0].VerifyPassword([]byte(r.Form.Get("password"))); ok {

		// Generate token
		token, err := ctx.DB.GenerateToken(results[0], ctx.Conf.Security)
		if err != nil {
			Error(w, http.StatusInternalServerError, err)
		} else {
			Json(w, http.StatusOK, base64.RawURLEncoding.EncodeToString(token))
		}
	} else {
		Error(w, http.StatusUnauthorized, errors.New("invalid password"))
	}
}
