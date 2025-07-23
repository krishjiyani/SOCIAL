package api

import (
	//"github.com/krishjiyani/SOCIAL/internal/store"
	"krishjiyani/SOCIAL/internal/store"
	"net/http"
"strconv"
	"github.com/go-chi/chi/v5"
	
)

func (app *Application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam("userID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//ctx := r.Context()

	user, err := app.getUser(r.Context(), userID)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.badRequestResponse(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.InternalServerError(w, r, err)
	}
}
