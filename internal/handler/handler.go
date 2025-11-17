package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kastuell/gotodoapp/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", h.register)
			// auth.POST("/sign-in", h.signIn)
		}

		profile := api.Group("/profile", h.userIdentity)
		{
			profile.GET("/", h.getMe)
		}

		list := api.Group("/list", h.userIdentity)
		{
			// lists.POST("/", h.createList)
			// lists.GET("/", h.getAllLists)
			// lists.GET("/:id", h.getListById)
			// lists.PUT("/:id", h.updateList)
			// lists.DELETE("/:id", h.deleteList)

			todo := list.Group(":id/todo")
			{
				todo.POST("/", h.createItem)
				todo.GET("/", h.getAllItems)
			}
		}

		todo := api.Group("todo", h.userIdentity)
		{
			todo.GET("/:id", h.getItemById)
			todo.PUT("/:id", h.updateItem)
			todo.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
