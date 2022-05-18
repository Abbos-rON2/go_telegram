package handlers

import (
	"strings"
	"telegram_back/models"

	"github.com/gin-gonic/gin"
)

// import (
// 	"telegram_back/models"

// 	"github.com/gin-gonic/gin"
// )

// @Summary Create chat
// @Description Create chat
// @Tags chats
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body models.CreateMessageRequest true "data"
// @Param members query string true "data"
// @Param type query string true "data"
// @Success 200 {object} models.Response{data=string}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /chats [post]
func (h *handler) CreateChat(c *gin.Context) {
	var message models.CreateMessageRequest
	m := c.Query("members")
	chatType := c.Query("type")
	members := strings.Split(m, ",")
	if err := c.BindJSON(&message); err != nil {
		h.handleError(c, err)
		return
	}

	chatId, err := h.storage.CreateChat(c.Request.Context(), chatType, members, message)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, chatId)
}

// @Summary Get chat
// @Description Get chat
// @Tags chats
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} models.Response{data=models.Chat}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /chats/{id} [get]
func (h *handler) GetChat(c *gin.Context) {
	var chat models.Chat
	if err := h.storage.GetChat(c.Request.Context(), c.Param("id"), &chat); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, chat)
}

// @Summary Get all chats
// @Description Get all chats
// @Tags chats
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_id query string true "user_id"
// @Success 200 {object} models.Response{data=[]models.Chat}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /chats [get]
func (h *handler) GetAllChats(c *gin.Context) {
	userID := c.Query("user_id")
	var chats []models.Chat
	if err := h.storage.GetAllChats(c.Request.Context(), userID, &chats); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, chats)
}

// // @Summary Get user chats
// // @Description Get user chats
// // @Tags chats
// // @Accept  json
// // @Produce  json
// // @Security ApiKeyAuth
// // @Param user_id path string true "user_id"
// // @Success 200 {object} models.Response{data=[]models.Chat}
// // @Failure 400 {object} models.Response
// // @Failure 500 {object} models.Response
// // @Router /chats/user/{user_id} [get]
// func (h *handler) GetUserChats(c *gin.Context) {
// 	var chats []models.Chat
// 	if err := h.storage.Chat().GetUserChats(c.Request.Context(), c.Param("user_id"), &chats); err != nil {
// 		h.handleError(c, err)
// 		return
// 	}

// 	h.handleSuccess(c, chats)
// }

// // @Summary Add user to chat
// // @Description Add user to chat
// // @Tags chats
// // @Accept  json
// // @Produce  json
// // @Security ApiKeyAuth
// // @Param chat_id path string true "chat_id"
// // @Param user_id path string true "user_id"
// // @Success 200 {object} models.Response{data=string}
// // @Failure 400 {object} models.Response
// // @Failure 500 {object} models.Response
// // @Router /chats/{chat_id}/users/{user_id} [post]
// func (h *handler) AddUserToChat(c *gin.Context) {
// 	if err := h.storage.Chat().AddUserToChat(c.Request.Context(), c.Param("chat_id"), c.Param("user_id")); err != nil {
// 		h.handleError(c, err)
// 		return
// 	}

// 	h.handleSuccess(c, "user added to chat")
// }

// // @Summary Remove user from chat
// // @Description Remove user from chat
// // @Tags chats
// // @Accept  json
// // @Produce  json
// // @Security ApiKeyAuth
// // @Param chat_id path string true "chat_id"
// // @Param user_id path string true "user_id"
// // @Success 200 {object} models.Response{data=string}
// // @Failure 400 {object} models.Response
// // @Failure 500 {object} models.Response
// // @Router /chats/{chat_id}/users/{user_id} [delete]
// func (h *handler) RemoveUserFromChat(c *gin.Context) {
// 	if err := h.storage.Chat().RemoveUserFromChat(c.Request.Context(), c.Param("chat_id"), c.Param("user_id")); err != nil {
// 		h.handleError(c, err)
// 		return
// 	}

// 	h.handleSuccess(c, "user removed from chat")
// }

// func (h *handler) GetChatUsers(c *gin.Context) {
// 	var users []models.UserDTO
// 	if err := h.storage.Chat().GetChatUsers(c.Request.Context(), c.Param("chat_id"), &users); err != nil {
// 		h.handleError(c, err)
// 		return
// 	}

// 	h.handleSuccess(c, users)
// }
