package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// If the use does not exist, it will be created,
// and an identifier is returned.
// If the user already exists, the identifier is returned.
func (rt *_router) createSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	username := r.FormValue("username")

	id, err := rt.db.GetUserId(username)
	if err != nil {
		// Create new user
		id, err = rt.db.CreateId("Users")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = rt.db.AddUser(id, username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx.Logger.Info("User ", username, " created")
	}
	ctx.Logger.Info("User ", username, " logged in")
	user := User{
		Username: username,
		UserId:   id,
	}
	_ = json.NewEncoder(w).Encode(user)

}
