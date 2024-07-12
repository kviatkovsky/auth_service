package user

import "context"

type User struct{
	ID        int64  `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Firstname string `json:"first_name" db:"first_name"`
	Lastname  string `json:"last_name" db:"last_name"`
	City      string `json:"city" db:"city"`
	School    string `json:"school" db:"school"`
}

type GetProfileRes struct {
	User
}

type Repository interface {
	GetProfiles(ctx context.Context, username string) ([]User, error)
}

type Service interface {
	GetProfiles(c context.Context, username string) ([]GetProfileRes, error)
}