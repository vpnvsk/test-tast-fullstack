package handler

import (
	"backend/pkg/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository *repository.Repository
}

func NewHandler(repository *repository.Repository) *Handler {
	return &Handler{repository: repository}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/", h.createUser)
			user.GET("/", h.getAllUsers)
			user.GET("/:id", h.getUserByID)
			user.PUT("/:id", h.updateUser)
			user.DELETE("/:id", h.deleteUser)
		}
	}
	return router
}
