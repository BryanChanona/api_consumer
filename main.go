package main

import (
	dependenciesUser "github.com/BryanChanona/api_consumer/src/User/infraestructure/dependencies"
	routesUser "github.com/BryanChanona/api_consumer/src/User/infraestructure/routes"
	"github.com/BryanChanona/api_consumer/src/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	dependenciesUser.Init()
	 
	 
	 r:= gin.Default()
	 helpers.InitCORS(r)
	routesUser.Routes(r)
	 r.Run(":8081")
}