package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
	}

	Users interface {
		Create(context.Context, *User) error
	}
}

func NewStorage(Db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{Db},
		Users: &UsersStore{Db},
	}
}

// app.store.Posts.Update()
// app.store.Users.
