package user

import (
	"context"
	"database/sql"
	"fmt"
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

func (r *repository) GetProfiles(ctx context.Context, username string) ([]User, error) {
	var rows *sql.Rows
	var err error
	query := `
		SELECT user.id, user.username, user_profile.first_name, user_profile.last_name, user_profile.city, user_data.school
		FROM user
		INNER JOIN user_profile
		ON user.id = user_profile.user_id
		INNER JOIN user_data
		ON user.id = user_data.user_id
	`

	if username != "" {
		whereClouse := "WHERE username = ?"
		query := fmt.Sprintf("%s %s", query, whereClouse)
		
		rows, err = r.db.QueryContext(ctx, query, username)
	} else {
		rows, err = r.db.QueryContext(ctx, query)
	}	

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Firstname, &u.Lastname, &u.City, &u.School); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) GetAuthByApiKey(ctx context.Context, apiKey string) (*AuthData, error) {
	auth := AuthData{}
	query := "SELECT id, `api-key` FROM auth WHERE `api-key` = ?"
	err := r.db.QueryRowContext(ctx, query, apiKey).Scan(&auth.ID, &auth.ApiKey)
	if err != nil {
		return &AuthData{}, err
	}

	return &auth, nil
}