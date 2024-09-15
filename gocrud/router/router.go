package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gocrud/controller"
	"net/http"
)

func NewRouter(userController *controller.UserController) *gin.Engine {
	service := gin.Default()

	// Configuração básica do CORS
	service.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // Permitir todas as origens
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},          // Permitir métodos específicos
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Permitir cabeçalhos específicos
		AllowCredentials: true,
	}))

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	userRouter := router.Group("/user")
	userRouter.GET("", userController.FindAll)
	userRouter.GET("/:userId", userController.FindById)
	userRouter.POST("", userController.Create)
	userRouter.PATCH("/:userId", userController.Update)
	userRouter.DELETE("/:userId", userController.Delete)

	return service
}
