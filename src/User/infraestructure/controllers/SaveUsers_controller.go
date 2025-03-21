package controllers

import (
	"net/http"

	"github.com/BryanChanona/api_consumer/src/User/application"
	"github.com/BryanChanona/api_consumer/src/User/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	useCase *application.RegisterUserUseCase
}

func NewUserController(useCase *application.RegisterUserUseCase) *UserController{
	return &UserController{useCase: useCase}
}

func (controller *UserController) Execute(ctx *gin.Context){
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	 err := controller.useCase.Execute(user)

	 if err != nil{
		ctx.JSON(500, gin.H{"error": err.Error()})
	}else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Successfully registered user"})
	}

}