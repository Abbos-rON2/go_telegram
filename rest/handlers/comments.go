package handlers

// import (
// 	"telegram_back/models"

// 	"github.com/gin-gonic/gin"
// )

// // @Summary Create comment
// // @Description Create comment
// // @Tags comments
// // @Accept  json
// // @Produce  json
// // @Security ApiKeyAuth
// // @Param data body models.CreateCommentRequest true "data"
// // @Success 200 {object} models.Response{data=string}
// // @Failure 400 {object} models.Response
// // @Failure 500 {object} models.Response
// // @Router /comments [post]
// func (h *handler) CreateComment(c *gin.Context) {
// 	var comment models.CreateCommentRequest
// 	if err := c.ShouldBindJSON(&comment); err != nil {
// 		return
// 	}
// 	comment.AuthorID = c.GetInt("user_id")

// 	if err := h.storage.Comment().CreateComment(c, comment); err != nil {
// 		h.handleError(c, err)
// 		return
// 	}

// 	h.handleSuccess(c, "comment created")
// }
