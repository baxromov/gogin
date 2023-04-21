package routes

import (
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"gogif/controllers"
	"gogif/repositories"
	"gogif/services"
)

func Initialize(database *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repositories.NewUserRepository(database)

	userService := services.NewUserService(userRepo)

	userController := controllers.NewUserController(userService)

	router.GET("/users", userController.List)

	return router
}
