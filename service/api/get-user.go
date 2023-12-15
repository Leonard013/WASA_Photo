package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Upload a photo.
// The user must be already logged in.
func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	token := r.URL.Query().Get("token")
	err := rt.db.CheckToken(token)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	username := ps.ByName("username")
	id, err := rt.db.GetUserId(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := User{
		UserId:   id,
		Username: username,
	}
	json.NewEncoder(w).Encode(user)

}
