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
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	var like phtoId_userId
	_ = json.NewDecoder(r.Body).Decode(&like)
	_ = r.Body.Close()
	photoId := like.PhotoId
	userId := like.UserId

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

	id, err := rt.db.PutLike(photoId, userId)
	if !errors.Is(err, nil) {
		if err.Error() == "already liked" {
			w.WriteHeader(http.StatusForbidden)
		} else if err.Error() == "photo not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	liker, err := rt.db.GetUsername(userId)
	if !errors.Is(err, nil) && errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if !errors.Is(err, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Logger.Info(liker, " liked ", photoId)
	_ = json.NewEncoder(w).Encode(Like{
		LikeId:  id,
		PhotoId: photoId,
		Author:  userId,
	})
}
