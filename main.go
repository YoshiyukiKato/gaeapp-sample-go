package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"cloud.google.com/go/datastore"
)

func main() {
	loadEnv()
	router := gin.Default()
	v1 := router.Group("/v1"){
		v1.GET("/gimmick", getGimmicks)
		v1.POST("/gimmick/new", newGimmick)
		v1.GET("/gimmick/:id", getGimmick)
		v1.POST("/gimmick/:id", setGimmick)
	}
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and server on 0.0.0.0:8080
}

func loadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//https://cloud.google.com/appengine/docs/go/datastore/reference
func client(){
	projectId := os.Getenv("PROJECT_ID")
	
	//client <-> context (?)
	client, err := google.DefaultClient(
		oauth2.NoContext,
		"https://www.googleapis.com/auth/devstorage.read_only")
	
	if err != nil {
		return nil, err
	}

	// Create the Google Cloud Storage service
	query, err := datastore.NewQuery()
	if err != nil {
		return nil, err
	}

	// Execute the request
	result, err := query.Run(client)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getGimmicks(){

}