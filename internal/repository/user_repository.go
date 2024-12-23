package repository

import (
	"go-gin-framework/internal/base/data"
	"go-gin-framework/internal/model"
)

type UserRepository struct {
	data *data.Data
}

func NewUserRepository(data *data.Data) *UserRepository {
	return &UserRepository{
		data: data,
	}
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	_, err := r.data.DB.Where("username = ?", username).Get(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
