package repository

import (
	"context"

	"github.com/facily-tech/go-scaffold/pkg/domains/user/model"
)

type UserRepositoryI interface {
	// Queries is a "Readeble" interface responsible to read data from source
	Querier

	// Execer is a "Writable" interface responsible for write data into source
	Execer
}

type Querier interface {
	FindByID(context.Context, string) string
}

type Execer interface {
	CreateUser(model.User) (model.User, error)
}
