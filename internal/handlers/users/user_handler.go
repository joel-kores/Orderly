package users

import (
	"Orderly/internal/models"
	"Orderly/internal/services/users"
)

type UserHandler struct {
	UserService users.UserService
}

func NewUserHandler(userService users.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// GetAllUsers godoc
//
//	@Summary		Get all users
//	@Description	Get all users
//	@Tags			users
//	@Security		OAuth2Password
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.User
//	@Router			/users [get]
func (h *UserHandler) GetAllUsers() ([]models.User, error) {
	allUsers, err := h.UserService.GetAll()
	if err != nil {
		return nil, err
	}

	var response []models.User
	for _, user := range allUsers {
		response = append(response, models.User{
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		})
	}

	return response, nil
}
