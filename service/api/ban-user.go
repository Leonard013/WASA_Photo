package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Upload a photo.
// The user must be already logged in.
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	token := r.URL.Query().Get("token")
	err := rt.db.CheckToken(token)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.ParseForm()
	username := r.FormValue("username")
	bannerId := r.FormValue("userId")
	bannedId, err := rt.db.GetUserId(username)
	if err != nil && err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.IfBanned(bannedId, bannerId) // check if it is blocked
	if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if err == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.Ban(bannerId, bannedId) // ban
	if err != nil {
		if err.Error() == "already following" {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	name, err := rt.db.GetUsername(bannerId)
	if err != nil && err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Info("User ", name, " banned ", username)
	user := User{
		Username: username,
		UserId:   bannedId,
	}
	_ = json.NewEncoder(w).Encode(user)
}
