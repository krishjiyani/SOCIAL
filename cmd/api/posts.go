package api

import (
	"context"
	"errors"
	//"fmt"
	"krishjiyani/SOCIAL/internal/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
)
type postKey string
const postCtx postKey = "post"

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

func (app *Application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//user := getUserFromContext(r)
	// ... (rest of the code remains the same)
	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		//change after auth
		UserID: 1,
	}

	ctx := r.Context()

	if err := app.Store.Posts.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
func (app *Application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	comments, err := app.store.Comments.GetByPostID(r.Context(), post.ID)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	post.Comments = comments

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.InternalServerError(w, r, err)
		return
// func (app *Application) getPostHandler(w http.ResponseWriter, r *http.Request) {
// 	idParam := chi.URLParam(r, "postID")
// 	id, err := strconv.ParseInt(idParam, 10, 64)
// 	if err != nil {
// 		app.InternalServerError(w, r, err)
// 		return
// 	}

// 	ctx := r.Context()

// 	post, err := app.store.Posts.GetByID(ctx, id)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, store.ErrNotFound):
// 			app.notFoundResponse(w, r, err)
// 		default:
// 			writeJSONError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}

// 	comments, err := app.store.Comments.GetByPostID(ctx, id)
// 	if err != nil {
// 		app.InternalServerError(w, r, err)
// 		return
// 	}

// 	post.Comments = comments

// 	if err := writeJSON(w, http.StatusOK, post); err != nil {
// 		app.InternalServerError(w, r, err)
// 		return
	}

}

func (app *Application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Posts.Delete(ctx, id); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.InternalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

 
type  UpdatePostPayload struct{
	Title *string `json:"title" validate:"omitempty,max=100"`
	Content *string `json:"content" validate:"omitempty,max=1000"`
}

func (app *Application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	var payload UpdatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Content != nil {
		post.Content = *payload.Content
	}
	if payload.Title != nil {
		post.Title = *payload.Title
	}

	ctx := r.Context()

	if err := app.updatePost(ctx, post); err != nil {
		app.InternalServerError(w, r, err)
	}

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.InternalServerError(w, r, err)
	}
}
	func (app *Application) postsContextMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idParam := chi.URLParam(r, "postID")
			id, err := strconv.ParseInt(idParam, 10, 64)
			if err != nil {
				app.InternalServerError(w, r, err)
				return
			}
	
			ctx := r.Context()
	
			post, err := app.store.Posts.GetByID(ctx, id)
			if err != nil {
				switch {
				case errors.Is(err, store.ErrNotFound):
					app.notFoundResponse(w, r, err)
				default:
					app.InternalServerError(w, r, err)
				}
				return
			}
	
			ctx = context.WithValue(ctx, postCtx, post)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	
	func getPostFromCtx(r *http.Request) *store.Post {
		post, _ := r.Context().Value(postCtx).(*store.Post)
		return post
	}
	
	func (app *Application) updatePost(ctx context.Context, post *store.Post) error {
		if err := app.store.Posts.Update(ctx, post); err != nil {
			return err
		}
	
		// app.cacheStorage.Users.Delete(ctx, post.UserID)
		return nil
	}

