package handler

import (
	"Zefirnutiy/simpleTodoApp/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()
	
	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.singUp)
		auth.POST("/sing-in",  h.singIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			todos := lists.Group("/todos")
			{
				todos.POST("/", h.createTodo)
				todos.GET("/", h.getAllTodos)
				todos.GET("/:todos_id", h.getTodoById)
				todos.PUT("/:todos_id", h.updateTodo)
				todos.DELETE("/:todos_id", h.deleteTodo)
			}
		}
	}

	return router  
}