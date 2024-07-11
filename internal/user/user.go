package user

import "context"

type User struct{
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

type GetProfileRes struct {
	ID          string `json:"id" db:"id"`
	Username    string `json:"username" db:"username"`
}

type Repository interface {
	GetProfile(ctx context.Context, username string) (*User, error)
}

type Service interface {
	GetProfile(c context.Context, username string) (*GetProfileRes, error)
}