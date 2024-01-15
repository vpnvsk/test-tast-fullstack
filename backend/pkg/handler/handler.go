package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vpnvsk/test-tast-fullstack/tree/main/backend/docs"
	"github.com/vpnvsk/test-tast-fullstack/tree/main/backend/pkg/repository"
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

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
