package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Upload a photo.
// The user must be already logged in.
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	_ = r.ParseForm()
	token := r.Header.Get("Authorization")
	err := rt.db.CheckToken(token)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photoId := ps.ByName("photoId") // the user to unban
	userId := r.Header.Get("userId") // the user who wants to unban

	authorId, err := rt.db.GetAuthorId(photoId)
	if !errors.Is(err, nil) {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = rt.db.IfBanned(authorId, userId) // check if it is blocked
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if errors.Is(err, nil) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.Unlike(photoId, userId)
	if !errors.Is(err, nil) {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	name, err := rt.db.GetUsername(userId)
	if !errors.Is(err, nil) && errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if !errors.Is(err, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Logger.Info("User ", name, " unlikes photo  ", photoId)
	_ = json.NewEncoder(w).Encode("Photo successfully unliked")
}
