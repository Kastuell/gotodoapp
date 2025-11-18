package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kastuell/gotodoapp/internal/auth"
	"github.com/kastuell/gotodoapp/internal/service"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
		h.initTodoRoutes(v1)
		h.initUserRoutes(v1)
	}
}
