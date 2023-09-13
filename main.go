package main

import (
	"log"
	"task/constant"
	"task/route"
	"github.com/gin-gonic/gin"
)



func main(){
	router:=gin.Default()
	route.Route(router)
	log.Fatal(router.Run(constant.Port))
}
