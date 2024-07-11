package user

import (
	"context"
	"strconv"
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

func (s *service) GetProfile(c context.Context, username string) (*GetProfileRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	r, err := s.Repository.GetProfile(ctx, username)
	if err != nil {
		return nil, err
	}

	res := &GetProfileRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
	}

	return res, nil
}