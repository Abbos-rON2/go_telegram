package handlers

import (
	"telegram_back/models"

	"github.com/gin-gonic/gin"
)

func makeUserDTO(u models.User) models.UserDTO {
	return models.UserDTO{
		ID:        u.ID,
		Phone:     u.Phone,
		Name:      u.Name,
		Bio:       u.Bio,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}

// @Summary Get a user
// @Description Get a user
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.Response{data=models.UserDTO}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /users/{id} [get]
func (h *handler) GetUser(c *gin.Context) {
	var user models.User

	err := h.storage.GetUser(c, c.Param("id"), &user)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccess(c, makeUserDTO(user))
}

// @Summary Get user by phone
// @Description Get user by phone
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param phone path string true "User phone"
// @Success 200 {object} models.Response{data=models.UserDTO}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /users/phone/{phone} [get]
func (h *handler) GetUserByPhone(c *gin.Context) {
	var user models.User

	err := h.storage.GetUserByPhone(c, c.Param("phone"), &user)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccess(c, makeUserDTO(user))
}

// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Response{data=[]models.UserDTO}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /users [get]
func (h *handler) GetAllUsers(c *gin.Context) {
	var users []models.User
	var usersDTO []models.UserDTO

	err := h.storage.GetAllUsers(c, &users)
	if err != nil {
		h.handleError(c, err)
		return
	}

	for _, v := range users {
		usersDTO = append(usersDTO, makeUserDTO(v))
	}

	h.handleSuccess(c, usersDTO)
}
