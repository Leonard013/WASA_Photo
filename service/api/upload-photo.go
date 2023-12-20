package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Upload a photo.
// The user must be already logged in.
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "multipart/form-data")
	token := r.URL.Query().Get("token")
	err := rt.db.CheckToken(token)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	title := r.FormValue("title")
	if len(title) > 30 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId := r.FormValue("userId")
	if _, err = rt.db.GetUsername(userId); err != nil && err == sql.ErrNoRows { // Check if the user is logged in
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	const maxFileSize = 20 << 20
	if header.Size > maxFileSize {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if header.Header.Get("Content-Type") != "image/png" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := rt.db.CreatePhotoId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	path, err := SavePhoto(file, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t := time.Now().Format("2006-01-02 15:04:05")

	err = rt.db.UploadPhoto(id, title, path, userId, t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Logger.Info("Photo uploaded")
	_ = json.NewEncoder(w).Encode(Photo{
		PhotoId:   id,
		Title:     title,
		PhotoPath: path,
		Date:      t,
		Author:    userId,
	})

}
