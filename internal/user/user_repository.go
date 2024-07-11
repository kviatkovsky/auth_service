package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) GetProfile(ctx context.Context, username string) (*User, error) {
	u := User{}
	query := "SELECT id, username FROM user WHERE username = ?"
	err := r.db.QueryRowContext(ctx, query, "test").Scan(&u.ID, &u.Username)
	if err != nil {
		return &User{}, err
	}

	return &u, nil
}