package handlers

import (
	"context"
	"errors"
	"net/http"
	"service/auth/app/users"
	usersdtos "service/auth/app/users/dtos"
	httperrors "service/auth/communications/http/misc/errors"
	httprequests "service/auth/communications/http/requests"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userAuthtenticator UserAuthtenticator
}

type UserAuthtenticator interface {
	Login(context.Context, usersdtos.LoginRequest) (string, error)
}

func NewAuthHandler(UserAuthtenticator UserAuthtenticator) *AuthHandler {
	return &AuthHandler{
		userAuthtenticator: UserAuthtenticator,
	}
}

func (h *AuthHandler) Url() string {
	return "v1/auth"
}

func (h *AuthHandler) Init(r *gin.RouterGroup) {
	r.POST("/login", h.login)
}

func (h *AuthHandler) login(c *gin.Context) {
	var req httprequests.LoginRequest
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	token, err := h.userAuthtenticator.Login(c, usersdtos.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		if errors.Is(err, users.ErrInvalidCredentials) {
			httperrors.BadRequestDetails(c, err)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
