package account

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}

func (s service) CreateUser(ctx context.Context, email, password string) (id string, err error) {
	logger := log.With(s.logger, "method", "CreateUser")

	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	id = uuid.String()

	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	if err := s.repository.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id)

	return id, nil

}

func (s service) GetUser(ctx context.Context, id string) (email string, err error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err = s.repository.GetUser(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get user", id)

	return email, nil
}
