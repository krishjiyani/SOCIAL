package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)
var (
	ErrNotFound = errors.New("resource not found")
	QueryTimeoutDuration = time.Second * 5 
)

type Storage struct {
	Posts interface {
		GetByID(context.Context, int64)(*Post, error)
		Create(context.Context, *Post) error
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
	}

	Users interface {
		GetByID(context.Context, int64) (*User, error)
		Create(context.Context, *User) error
	}
	Comments interface{
		Create(context.Context, *Comment) error
		GetByPostID( context.Context,  int64) ([]Comment, error)
	}
}

func NewStorage(Db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{Db},
		Users: &UsersStore{Db},
		Comments: &CommentStore{Db},
	}
}

// app.store.Posts.Update()
// app.store.Users.
