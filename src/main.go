package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/test", test)
		//v1.GET("/gimmick", getGimmicks)
		//v1.POST("/gimmick/new", newGimmick)
		//v1.GET("/gimmick/:id", getGimmick)
		//v1.POST("/gimmick/:id", setGimmick)
	}
	
	router.Run() // listen and server on 0.0.0.0:8080
}

func test(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "test!",
	})
}

//https://cloud.google.com/appengine/docs/go/datastore/reference
