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
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	streamId := ps.ByName("streamId")

	var stream []PhotoForStream

	if _, ok := images_stream[streamId]; !ok {
		dbStream, err := rt.db.GetStream(streamId)
		if !errors.Is(err, nil) {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusNotFound)
				return
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		stream = make([]PhotoForStream, len(dbStream))
		for i, v := range dbStream {
			stream[i] = PhotoForStream(v) // direct conversion is possible since the fields are the same
		}
		images_stream[streamId] = append(images_stream[streamId], stream...)
	}
	stream = images_stream[streamId]

	// take the first 20 elements and remove them from the slice
	if stream != nil {
		if len(stream) > 20 {
			stream = stream[:20]
			images_stream[streamId] = images_stream[streamId][20:]
		} else { // if there are less than 20 elements, take them all and remove them from the slice
			images_stream[streamId] = nil
		}
		_ = json.NewEncoder(w).Encode(stream)
	} else {
		_ = json.NewEncoder(w).Encode("There are no photos in the stream")
	}

}
