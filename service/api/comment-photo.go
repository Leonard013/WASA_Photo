package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Upload a photo.
// The user must be already logged in.
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "multipart/form-data")
	token := r.Header.Get("Authorization")
	err := rt.db.CheckToken(token)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photoId := r.FormValue("photoId")
	userId := r.FormValue("authorId")
	text := r.FormValue("text")

	if len(text) > 300 || len(text) < 1 {
		w.WriteHeader(http.StatusBadRequest)
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
	t := time.Now().Format("2006-01-02 15:04:05")
	id, err := rt.db.Comment(photoId, userId, text, t)
	if !errors.Is(err, nil) {
		if err.Error() == "photo does not exist" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	commenter, err := rt.db.GetUsername(userId)
	if !errors.Is(err, nil) && errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if !errors.Is(err, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Logger.Info(commenter, " commentes ", photoId)
	_ = json.NewEncoder(w).Encode(Comment{
		CommentId: id,
		PhotoId:   photoId,
		Author:    userId,
		Text:      text,
		Date:      t,
	})
}
