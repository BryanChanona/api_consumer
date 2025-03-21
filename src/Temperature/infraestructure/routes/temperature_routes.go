package routes

import (
	"github.com/BryanChanona/api_consumer/src/Temperature/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	routes := router.Group("/users")
	saveTemperature := dependencies.GetSaveTemperatureController().Execute
	
	routes.GET("/", saveTemperature )
}