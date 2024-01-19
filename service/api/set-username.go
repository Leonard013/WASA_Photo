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

	username := ps.ByName("username")
	userId := r.Header.Get("userId")
	u, err := rt.db.GetUserId(username)
	if !errors.Is(err, nil) && errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if !errors.Is(err, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userId != u {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	_ = r.ParseForm()
	username = r.FormValue("username")
	if len(username) > 20 || len(username) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.SetUsername(username, userId)
	if !errors.Is(err, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Logger.Infof("Username changed to %s", username)

	_ = json.NewEncoder(w).Encode(User{
		Username: username,
		UserId:   userId,
	})

}
