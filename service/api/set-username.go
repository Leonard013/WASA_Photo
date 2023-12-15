package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Upload a photo.
// The user must be already logged in.
func (rt *_router) setUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	token := r.URL.Query().Get("token")
	err := rt.db.CheckToken(token)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	username := ps.ByName("username")
	if len(username) > 20 || len(username) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId := r.URL.Query().Get("userId")

	if u, err := rt.db.GetUserId(username); err != nil { // check if the user has the right to change the username
		w.WriteHeader(http.StatusNotFound)
		return
	} else if u != userId {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	r.ParseForm()
	username = r.FormValue("username")
	err = rt.db.SetUsername(username, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Logger.Infof("Username changed to %s", username)

	_ = json.NewEncoder(w).Encode(User{
		Username: username,
		UserId:   userId,
	})

}
