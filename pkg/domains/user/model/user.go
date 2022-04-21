package model

import "gorm.io/gorm"

// A struct User é a struct raiz e é a única persistida no banco!
type User struct {
	gorm.Model
	ID       int `gorm:"primaryKey autoIncrement:true"`
	Name     string
	BirthDay string
}

// Structs que vem via request
type CreateUserRequest struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name" validate:"required"`
	BirthDay string `json:"birthday" validate:"required"`
}

type UpdateUserRequest struct {
	ID       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	BirthDay string `json:"birthday" validate:"required"`
}
type GetUserByIDRequest struct {
	ID int `json:"id" validate:"required"`
}

type DeleteUserByIDRequest struct {
	ID string `json:"id" validate:"required"`
}

// Struct de resposta da request
type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	BirthDay string `json:"birthday"`
}

// Construtores para fazer a conversao para um user
func NewUserFromCreateRequest(user CreateUserRequest) User {
	return User{
		Name:     user.Name,
		BirthDay: user.BirthDay,
	}
}

func NewUserFromUpdateRequest(user UpdateUserRequest) User {
	return User{
		ID:       user.ID,
		Name:     user.Name,
		BirthDay: user.BirthDay,
	}
}

func NewUserResponse(user User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		BirthDay: user.BirthDay,
	}
}
