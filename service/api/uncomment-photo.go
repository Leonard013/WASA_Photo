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
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	commentId := ps.ByName("commentId")     // the user to unban
	photoId := r.URL.Query().Get("photoId") // the user to unban
	userId := r.URL.Query().Get("userId")

	authorId, err := rt.db.GetAuthorId(photoId)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = rt.db.IfBanned(authorId, userId) // check if it is blocked
	if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if err == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.Uncomment(commentId, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	name, err := rt.db.GetUsername(userId)
	if err != nil && err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Info("User ", name, " uncomments photo  ", photoId)
	_ = json.NewEncoder(w).Encode("Photo successfully uncommented")
}