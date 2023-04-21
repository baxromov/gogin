package controllers

import (
	gin "github.com/gin-gonic/gin"
	models "gogif/db"
	services "gogif/services"
)

type UserController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
}

type userController struct {
	service services.UserService
}

// Create implements UserController
func (*userController) Create(c *gin.Context) {
	panic("unimplemented")
}

// Delete implements UserController
func (*userController) Delete(c *gin.Context) {
	panic("unimplemented")
}

// GetByID implements UserController
func (*userController) GetByID(c *gin.Context) {
	panic("unimplemented")
}

// Update implements UserController
func (*userController) Update(c *gin.Context) {
	panic("unimplemented")
}

func NewUserController(userCon services.UserService) UserController {
	return &userController{service: userCon}
}

func (c *userController) List(ctx *gin.Context) {
	var users []models.User
	if err := c.service.List(&users); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, users)

}
