package route

import (
	"task/controller"

	"github.com/gin-gonic/gin"
)

func Route(r * gin.Engine){
	r.POST("/test",controller.Test)
	r.POST("/tokens",controller.Token)
	r.GET("/tokensgetting",controller.GetToken)
}