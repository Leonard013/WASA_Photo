package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// If the use does not exist, it will be created,
// and an identifier is returned.
// If the user already exists, the identifier is returned.
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	var login Username
	_ = json.NewDecoder(r.Body).Decode(&login)
	_ = r.Body.Close()
	username := login.Username

	id, err := rt.db.GetUserId(username)

	if !errors.Is(err, nil) {
		if len(username) > 20 || len(username) < 3 {
			w.WriteHeader(http.StatusBadRequest)
		}
		id, err = rt.db.AddUser(username)
		if !errors.Is(err, nil) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx.Logger.Info("User ", username, " created")
	}
	ctx.Logger.Info("User ", username, " logged in")

	info, err := rt.db.GetUserInfo(id)
	if !errors.Is(err, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	user := User{
		Username:  username,
		UserId:    id,
		Followers: info.Followers,
		Following: info.Following,
		Banned:    info.Banned,
		IsBanned:  info.IsBanned,
	}
	_ = json.NewEncoder(w).Encode(user)

}
