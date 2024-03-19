package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Upload a photo.
// The user must be already logged in.
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	var new_username Username
	_ = json.NewDecoder(r.Body).Decode(&new_username)
	_ = r.Body.Close()
	username := new_username.Username



	_ = r.ParseForm()
	token := r.Header.Get("Authorization")
	err := rt.db.CheckToken(token)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userId := r.Header.Get("userId")

	_, err = rt.db.GetUserId(username)

	


	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if errors.Is(err, nil) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if len(username) > 20 || len(username) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.SetUsername(username, userId)
	if !errors.Is(err, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Logger.Infof("Username changed from '%s' to '%s'",ps.ByName("username"), username)

	var user User
	user.Username = username
	user.UserId = userId

	_ = json.NewEncoder(w).Encode(user)

}
