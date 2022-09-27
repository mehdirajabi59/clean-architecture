package http

import (
	"clean/entity"
	"clean/user/repository"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	FindAll(c *gin.Context)
	Find(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type userHandler struct {
	userUsecase entity.UserUseCase
}

func NewUserHandler(uc entity.UserUseCase) UserHandler {
	return &userHandler{
		userUsecase: uc,
	}
}

func (uh *userHandler) FindAll(c *gin.Context) {

	users, err := uh.userUsecase.FindAll()

	if err != nil {
		c.JSON(500, gin.H{"message": "The server unavailable temprory"})
		return
	}

	c.JSON(200, users)
}

func (uh *userHandler) Find(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{"message": "The server unavailable temprory"})
		return
	}

	user, err := uh.userUsecase.Find(int64(userId))

	if err != nil {
		c.JSON(404, gin.H{"message": "The user not found"})
		return
	}

	c.JSON(200, user)
}

func (uh *userHandler) Create(c *gin.Context) {

	var u entity.User
	c.Bind(&u)

	if err := uh.userUsecase.Create(&u); err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"message": "The server unavailable temprory"})
		return
	}

	c.JSON(201, gin.H{"message": "The user created successfuly"})
}

func (uh *userHandler) Delete(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{"message": "The server unavailable temprory"})
		return
	}

	var u entity.User
	u.ID = int64(userId)

	err = uh.userUsecase.Delete(&u)

	if errors.Is(err, repository.ErrorUserNotFound) {
		c.JSON(404, gin.H{"message": fmt.Errorf("user_id %d is invalid: %w", userId, err).Error()})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"message": "The server unavailable temprory"})
		return
	}

	c.JSON(200, gin.H{"message": "The user deleted successfuly"})

}
