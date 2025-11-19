package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kastuell/gotodoapp/internal/auth"
	"github.com/kastuell/gotodoapp/internal/config"
	v1 "github.com/kastuell/gotodoapp/internal/http/v1"
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

func (h *Handler) InitRoutes(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	h.initAPI(router)

	return router

	// api := router.Group("/api")
	// {
	// 	auth := api.Group("/auth")
	// 	{
	// 		auth.POST("/register", h.register)
	// 		// auth.POST("/sign-in", h.signIn)
	// 	}

	// 	profile := api.Group("/profile", h.userIdentity)
	// 	{
	// 		profile.GET("/", h.getMe)
	// 	}

	// 	list := api.Group("/list", h.userIdentity)
	// 	{
	// 		// lists.POST("/", h.createList)
	// 		// lists.GET("/", h.getAllLists)
	// 		// lists.GET("/:id", h.getListById)
	// 		// lists.PUT("/:id", h.updateList)
	// 		// lists.DELETE("/:id", h.deleteList)

	// 		todo := list.Group(":id/todo")
	// 		{
	// 			todo.POST("/", h.createItem)
	// 			todo.GET("/", h.getAllItems)
	// 		}
	// 	}

	// 	todo := api.Group("todo", h.userIdentity)
	// 	{
	// 		todo.GET("/:id", h.getItemById)
	// 		todo.PUT("/:id", h.updateItem)
	// 		todo.DELETE("/:id", h.deleteItem)
	// 	}
	// }
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
