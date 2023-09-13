package main

import (
	"log"
	"net/http"
	"task/constant"
	"task/route"
	
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	route.Route(router)
	http.HandleFunc("/", token.handler)
	http.ListenAndServe(":4000", nil)
	log.Fatal(router.Run(constant.Port))
}
