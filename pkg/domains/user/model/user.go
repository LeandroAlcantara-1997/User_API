package model

import "gorm.io/gorm"

// É a única persistida no banco!
type User struct {
	gorm.Model
	ID       int `gorm:"primaryKey autoIncrement:true"`
	Name     string
	BirthDay string
}

// Structs que vem via request
type CreateUserRequest struct {
	ID       int    `json:"id" `
	Name     string `json:"name" validate:"required"`
	BirthDay string `json:"birthday"`
}

type UpdateUserRequest struct {
	ID       int
	Name     string
	BirthDay string
}
type GetUserByIDRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeleteByIDRequest struct {
	ID string `json:"id"`
}

func NewUserFromCreate(user CreateUserRequest) User {
	return User{
		Name:     user.Name,
		BirthDay: user.BirthDay,
	}
}

func NewUserFromUpdate(user UpdateUserRequest) User {
	return User{
		Name:     user.Name,
		BirthDay: user.BirthDay,
	}
}
