package store

import (
	"context"
	"database/sql"
	
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"` //we wont return the pass wheneverwe marshall or unmarshall
	CreatedAt string `json:"created_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, userID int64) (*User, error) {
 `  SELECT id, username, email, password, created_at
	FROM users
	WHERE id = $1
`

	ctx,cancel :=context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		user.ID
		).Scan(
		&user.ID,
		user.Username,
		user.Password,
		user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		switch err {
        case sql.ErrNoRows:
			return nil, ErrNotFound
			default:
				return nil, err		
		}
	

	}

	return user, nil
}



