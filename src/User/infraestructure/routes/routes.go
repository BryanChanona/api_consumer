package routes

import (

	"github.com/BryanChanona/api_consumer/src/User/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/users")

	saveController := dependencies.GetSaveUserController().Execute
	routes.POST("/",saveController)
}