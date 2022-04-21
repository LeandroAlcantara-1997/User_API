package repository

import (
	"context"

	userErr "github.com/facily-tech/go-scaffold/pkg/domains/user/error"
	"github.com/facily-tech/go-scaffold/pkg/domains/user/model"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	client *gorm.DB
}

func NewRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		client: db,
	}
}

func (p *PostgresRepository) CreateUser(user model.User) (model.User, error) {
	if err := p.client.AutoMigrate(&user); err != nil {
		return user, err
	}
	result := p.client.Create(&user)
	if result.RowsAffected == 0 {
		return user, userErr.ErrEmptyRepository
	}
	return user, nil
}
func (p *PostgresRepository) FindByID(ctx context.Context, id string) string {
	return ""
}
