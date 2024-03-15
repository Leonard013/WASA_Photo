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
func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	id, err := rt.db.GetUserId(username)
	if !errors.Is(err, nil) && errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if !errors.Is(err, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bannerId := r.Header.Get("userId")
	err = rt.db.IfBanned(id, bannerId) // check if it is blocked
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	} else if errors.Is(err, nil) {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Error(http.StatusForbidden)
		return
	}

	info, err := rt.db.GetUserInfo(id)
	if !errors.Is(err, nil) && errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if !errors.Is(err, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	

	user := User{
		UserId:   id,
		Username: username,
		Followers: info.Followers,
		Following: info.Following,
		Banned: info.Banned,
		IsBanned: info.IsBanned,
	}

	_ = json.NewEncoder(w).Encode(user)

}
