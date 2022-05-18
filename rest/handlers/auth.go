package handlers

import (
	"strings"
	"time"

	"telegram_back/errs"
	"telegram_back/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// @Summary Login
// @Description Login
// @Tags auth
// @Accept  json
// @Produce  json
// @Param data body models.Login true "data"
// @Success 200 {object} models.Response{data=models.User}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /login [post]
func (h *handler) Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		return
	}

	dbUser := models.User{}
	if err := h.storage.GetUserByPhone(c, user.Phone, &dbUser); err != nil {
		h.handleError(c, err)
		return
	}

	if user.Password != dbUser.Password {
		h.handleError(c, errs.ErrInvalidPassword)
		return
	}

	// create jwt token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Token{
		UserID: dbUser.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.cfg.JwtSecret))
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.Header("Authorization", "Bearer "+tokenString)

	h.handleSuccess(c, dbUser)
}

// @Summary Register
// @Description Register
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.CreateUserRequest true "User"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /register [post]
func (h *handler) CreateUser(c *gin.Context) {
	var user models.CreateUserRequest

	if err := c.ShouldBind(&user); err != nil {
		return
	}

	if err := h.storage.CreateUser(c, user); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, "user created")
}

func (h *handler) AuthMiddleware(c *gin.Context) {
	var claims = &models.Token{}

	strArr := strings.Split(c.Request.Header.Get("Authorization"), " ")
	if len(strArr) != 2 {
		h.handleError(c, errs.ErrInvalidToken)
		c.Abort()
		return
	}

	token := strArr[1]
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.cfg.JwtSecret), nil
	})
	if err != nil {
		h.handleError(c, errs.ErrUnauthorized)
		c.Abort()
		return
	}

	if !tkn.Valid {
		h.handleError(c, errs.ErrInvalidToken)
		c.Abort()
		return
	}
	c.Set("user_id", claims.UserID)
	c.Next()
}

func (h *handler) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
