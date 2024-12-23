package service

import (
	"errors"
	"go-gin-framework/internal/repository"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) ValidateCredentials(username, password string) (bool, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, errors.New("user not found")
	}

	// TODO: 实际应该比较加密后的密码
	return user.Password == password, nil
}
