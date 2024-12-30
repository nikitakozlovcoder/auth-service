package handlers

import (
	"context"
	"errors"
	"net/http"
	"service/auth/app/domain/requests"
	"service/auth/app/services"
	misc "service/auth/http/misc/errors"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userAuthtenticator UserAuthtenticator
}

type UserAuthtenticator interface {
	Login(context.Context, requests.LoginRequest) (string, error)
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
	var req requests.LoginRequest
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	token, err := h.userAuthtenticator.Login(c, req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			misc.BadRequestDetails(c, err)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
