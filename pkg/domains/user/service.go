package user

import (
	"context"
	"fmt"

	"github.com/facily-tech/go-core/log"

	"github.com/facily-tech/go-scaffold/pkg/domains/user/model"
	"github.com/facily-tech/go-scaffold/pkg/domains/user/repository"
)

type ServiceI interface {
	FindUserByID(context.Context, string) (model.CreateUserRequest, error)
	CreateUser(context.Context, model.CreateUserRequest) (model.User, error)
}

type Service struct {
	repository repository.UserRepositoryI
	log        log.Logger
}

func NewService(repository repository.UserRepositoryI, log log.Logger) (*Service, error) {
	if repository == nil {
		return nil, fmt.Errorf("Repositorio inesistente")
	}
	return &Service{
		repository: repository,
		log:        log,
	}, nil
}

func (s *Service) FindUserByID(ctx context.Context, id string) (model.CreateUserRequest, error) {
	return model.CreateUserRequest{}, nil
}

func (s *Service) CreateUser(ctx context.Context, user model.CreateUserRequest) (model.User, error) {
	userCreate, err := s.repository.CreateUser(model.NewUserFromCreate(user))
	if err != nil {
		return userCreate, err
	}
	return userCreate, nil
}
