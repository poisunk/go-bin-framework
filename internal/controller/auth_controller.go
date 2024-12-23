package controller

import (
	"go-gin-framework/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginReq struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request parameters"})
		return
	}

	valid, err := c.authService.ValidateCredentials(loginReq.Username, loginReq.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if !valid {
		ctx.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Login successful",
		"token":   "dummy-token", // TODO: 实现 JWT
	})
}
