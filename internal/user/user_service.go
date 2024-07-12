package user

import (
	"context"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (s *service) GetProfiles(c context.Context, username string) ([]GetProfileRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	profiles, err := s.Repository.GetProfiles(ctx, username)
	if err != nil {
		return nil, err
	}

	res := make([]GetProfileRes, len(profiles))

	for i, user := range profiles {
		res[i] = GetProfileRes{
			User{
				ID:        user.ID,
				Username:  user.Username,
				Firstname: user.Firstname,
				Lastname:  user.Lastname,
				City:      user.City,
				School:    user.School,
			},
		}
	}


	return res, nil
}

func(s *service) GetAuthByApiKey(c context.Context, apiKey string) (*AuthData, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	authData, err := s.Repository.GetAuthByApiKey(ctx, apiKey)
	if err != nil {
		return &AuthData{}, err
	}

	return authData, nil
}