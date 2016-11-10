package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	router := gin.New()
	v1 := router.Group("/v1")
	{
		v1.GET("/test", test)
		//v1.GET("/gimmick", getGimmicks)
		//v1.POST("/gimmick/new", newGimmick)
		//v1.GET("/gimmick/:id", getGimmick)
		//v1.POST("/gimmick/:id", setGimmick)
	}
	
	http.Handle("/", router)
}

func test(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "test!",
	})
}

//https://cloud.google.com/appengine/docs/go/datastore/reference
