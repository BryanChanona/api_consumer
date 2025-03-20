package main

import (
	"github.com/BryanChanona/api_consumer/src/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	 helpers.ConnMySQL()
	 
	 r:= gin.Default()
	 helpers.InitCORS(r)
	 r.Run(":8081")
}