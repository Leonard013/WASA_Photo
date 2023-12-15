package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Upload a photo.
// The user must be already logged in.
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	token := r.URL.Query().Get("token")
	err := rt.db.CheckToken(token)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	username := ps.ByName("username") // the user to unban
	unbanId, err := rt.db.GetUserId(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bannerId := r.URL.Query().Get("userId")
	err = rt.db.IfBanned(unbanId, bannerId) // check if it is blocked
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.Unban(bannerId, unbanId)
	if err != nil {
		if err.Error() == "already unfollowed" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	name, err := rt.db.GetUsername(bannerId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ctx.Logger.Info("User ", name, " unbanned ", username)
	_ = json.NewEncoder(w).Encode("User succesfully unbanned")
}
