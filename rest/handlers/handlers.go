package handlers

import (
	"context"
	"net/http"

	"telegram_back/config"
	"telegram_back/models"

	"github.com/gin-gonic/gin"
)

type UserI interface {
	CreateUser(ctx context.Context, user models.CreateUserRequest) error
	GetUser(ctx context.Context, id string, user *models.User) error
	GetUserByPhone(ctx context.Context, phone string, user *models.User) error
	GetAllUsers(ctx context.Context, users *[]models.User) error

	CreateMessage(ctx context.Context, message models.CreateMessageRequest) error
	CreateChat(ctx context.Context, chatType string, members []string, message models.CreateMessageRequest) (chatID string, err error)
	GetChat(ctx context.Context, chatID string, chat *models.Chat) error
	GetAllChats(ctx context.Context, userID string, chats *[]models.Chat) error
}
type handler struct {
	cfg     config.Config
	storage UserI
}

func NewHandler(cfg config.Config, storage UserI) *handler {
	return &handler{
		cfg:     cfg,
		storage: storage,
	}
}

func (h *handler) handleSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
func (h *handler) handleError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, models.Response{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	})
}
