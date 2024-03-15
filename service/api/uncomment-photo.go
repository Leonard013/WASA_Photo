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
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	// var info Uncomment
	// _ = json.NewDecoder(r.Body).Decode(&info)
	// _ = r.Body.Close()
	// photoId := info.PhotoId
	// userId := info.UserId
	// ctx.Logger.Info(photoId)
	// ctx.Logger.Info(userId)
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

	commentId := ps.ByName("commentId") // the user to unban
	photoId := r.Header.Get("photoId")  // the user to unban
	userId := r.Header.Get("userId")

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

	err = rt.db.Uncomment(commentId, userId)
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

	ctx.Logger.Info("User ", name, " uncomments photo  ", photoId)
	_ = json.NewEncoder(w).Encode("Photo successfully uncommented")
}
