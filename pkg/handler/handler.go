package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.signUp)
		auth.POST("/sing-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.POST("/:id", h.getListById)
			lists.GET("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group("/")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	return router
}
