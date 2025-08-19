package users

import (
	"Orderly/internal/models"
	"Orderly/internal/repositories/users"
)

type UserService struct {
	UserRepository users.UserRepository
}

func NewUserService(userRepo users.UserRepository) *UserService {
	return &UserService{UserRepository: userRepo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.UserRepository.GetAll()
}

func (s *UserService) Update(user *models.User) error {
	return s.UserRepository.Update(user)
}

func (s *UserService) Delete(id uint) error {
	return s.UserRepository.Delete(id)
}

func (s *UserService) GetCurrentUser(id uint) (*models.User, error) {
	return s.UserRepository.GetCurrentUser(id)
}
