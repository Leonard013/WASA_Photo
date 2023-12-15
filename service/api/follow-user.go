package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Upload a photo.
// The user must be already logged in.
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	token := r.URL.Query().Get("token")
	err := rt.db.CheckToken(token)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	r.ParseForm()
	username := r.FormValue("username")
	userId := r.FormValue("userId")
	followedId, err := rt.db.GetUserId(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = rt.db.IfBanned(followedId, userId) // check if it is blocked
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.Follow(userId, followedId)
	if err != nil {
		if err.Error() == "already following" {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	name, err := rt.db.GetUsername(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ctx.Logger.Info("User ", name, " followed ", username)
	user := User{
		Username: username,
		UserId:   followedId,
	}
	_ = json.NewEncoder(w).Encode(user)
}
