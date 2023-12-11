package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Upload a photo.
// The user must be already logged in.
func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	username := ps.ByName("username")
	id, err := rt.db.GetUser(username)
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
