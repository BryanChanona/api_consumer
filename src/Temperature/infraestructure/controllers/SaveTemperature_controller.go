package controllers

import (
	"net/http"

	"github.com/BryanChanona/api_consumer/src/Temperature/application"
	"github.com/BryanChanona/api_consumer/src/Temperature/domain"
	"github.com/gin-gonic/gin"
)

type SaveTemperatureController struct {
	useCase *application.SaveTemperatureUseCase
}

func NewSaveTemperatureController(useCase *application.SaveTemperatureUseCase ) *SaveTemperatureController{
	return &SaveTemperatureController{useCase: useCase}
}


func (controller *SaveTemperatureController) Execute(ctx *gin.Context){
	var temperature domain.Temperature

	if err := ctx.ShouldBindJSON(&temperature); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	 err := controller.useCase.Execute(temperature)

	 if err != nil{
		ctx.JSON(500, gin.H{"error": err.Error()})
	}else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Temperature saved"})
	}

	



}