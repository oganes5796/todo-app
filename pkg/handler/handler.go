package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/oganes5796/todo-app/pkg/service"
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
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		list := api.Group("/lists")
		{
			list.POST("/", h.createList)
			list.GET("/", h.getAllLists)
			list.GET("/:id", h.getListsById)
			list.PUT("/:id", h.updateList)
			list.DELETE("/:id", h.deleteList)

			items := list.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
