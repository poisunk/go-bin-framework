package router

import (
	"go-gin-framework/internal/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	authController *controller.AuthController
}

func NewRouter(authController *controller.AuthController) *Router {
	return &Router{
		authController: authController,
	}
}

func (r *Router) RegisterRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", r.authController.Login)
		}
	}
}
