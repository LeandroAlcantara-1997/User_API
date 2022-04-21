package repository

import (
	"context"

	userErr "github.com/facily-tech/go-scaffold/pkg/domains/user/error"
	"github.com/facily-tech/go-scaffold/pkg/domains/user/model"
	"gorm.io/gorm"
)

// A struct PostgresRepository deve implementar a interface UserRepositoryI
type PostgresRepository struct {
	client *gorm.DB
}

// A strcut PostgresRepository deve ter um construtor para ser injetada a dependência no construtor do serviço
func NewRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		client: db,
	}
}

// Cria um user
func (p *PostgresRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	tx := p.client.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return user, userErr.ErrNew
	}

	tx.Commit()
	return user, nil
}

// Encontra o user por id
func (p *PostgresRepository) FindUserByID(ctx context.Context, id int) (model.User, error) {
	var user model.User
	result := p.client.WithContext(ctx).Where(model.User{ID: id}).First(&user)
	if result.RowsAffected == 0 {
		return user, userErr.ErrNotFound
	}
	return user, nil
}

// Atualiza um user
func (p *PostgresRepository) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	tx := p.client.WithContext(ctx).Begin()
	if err := tx.WithContext(ctx).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		tx.WithContext(ctx).Rollback()
		return user, err
	}
	tx.WithContext(ctx).Commit()
	return user, nil
}
