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
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	_ = r.ParseForm()
	token := r.Header.Get("Authorization")
	err := rt.db.CheckToken(token)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photoId := ps.ByName("photoId") // the photo to delete
	userId := r.Header.Get("userId")

	path, err := rt.db.DeletePhoto(photoId, userId)
	if err != nil {
		if err.Error() == "already deleted" {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err.Error() == "not your photo" {
			w.WriteHeader(http.StatusForbidden)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	err = DeletePhoto(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Info("photo deleted")
	_ = json.NewEncoder(w).Encode("Photo successfully deleted")
}
