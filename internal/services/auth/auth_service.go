package services

import (
	"Orderly/internal/models"
	repositories "Orderly/internal/repositories/auth"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) RegisterOrLogin(email, name, phone string) (*models.User, error) {
	// Check if the user already exists
	user, err := s.UserRepo.FindByEmail(email)
	if err == nil {
		return user, nil
	}

	// User does not exist, create a new user
	newUser := &models.User{
		Name:  name,
		Email: email,
		Phone: phone,
	}
	err = s.UserRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
