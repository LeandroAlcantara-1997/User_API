package user

import (
	"context"

	"github.com/facily-tech/go-core/log"

	userErr "github.com/facily-tech/go-scaffold/pkg/domains/user/error"
	"github.com/facily-tech/go-scaffold/pkg/domains/user/model"
	"github.com/facily-tech/go-scaffold/pkg/domains/user/repository"
)

// Toda função de serviço deve ser declarada na interface ServiceI para se implementada
type ServiceI interface {
	FindUserByID(context.Context, int) (model.UserResponse, error)
	CreateUser(context.Context, model.CreateUserRequest) (model.UserResponse, error)
	UpdateUser(context.Context, model.UpdateUserRequest) (model.UserResponse, error)
	DeleteUserByID(context.Context, model.DeleteUserByIDRequest) error
}

// A struct Service deve implementar as funções da interface ServiceI
type Service struct {
	repository repository.UserRepositoryI
	log        log.Logger
}

// O Service deve ter um construtor, é chamado para iniciar os serviços no arquivo container.go
func NewService(repository repository.UserRepositoryI, log log.Logger) (*Service, error) {
	if repository == nil {
		return nil, userErr.ErrEmptyRepository
	}
	return &Service{
		repository: repository,
		log:        log,
	}, nil
}

// Busca um user por id
func (s *Service) FindUserByID(ctx context.Context, id int) (model.UserResponse, error) {
	user, err := s.repository.FindUserByID(ctx, id)
	if err != nil {
		return model.NewUserResponse(user), err
	}
	return model.NewUserResponse(user), nil
}

// Cria um user
func (s *Service) CreateUser(ctx context.Context, user model.CreateUserRequest) (model.UserResponse, error) {
	userCreate, err := s.repository.CreateUser(ctx, model.NewUserFromCreateRequest(user))
	if err != nil {
		return model.NewUserResponse(userCreate), err
	}
	return model.NewUserResponse(userCreate), nil
}

// Atualiza um user
func (s *Service) UpdateUser(ctx context.Context, user model.UpdateUserRequest) (model.UserResponse, error) {
	userUpdate, err := s.repository.UpdateUser(ctx, model.NewUserFromUpdateRequest(user))
	if err != nil {
		return model.NewUserResponse(userUpdate), err
	}
	return model.NewUserResponse(userUpdate), err
}

// Deleta user
func (s *Service) DeleteUserByID(ctx context.Context, userDelete model.DeleteUserByIDRequest) error {
	if err := s.repository.DeleteUser(ctx, userDelete.ID); err != nil {
		return err
	}

	return nil
}
