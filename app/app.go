package app

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"encoding/json"
)

func init() {
	router := gin.New()
	v1 := router.Group("/v1")
	{
		v1.GET("/gimmick", getGimmicks)
		v1.POST("/gimmick/new", newGimmick)
		v1.GET("/gimmick/:id", getGimmick)
		v1.POST("/gimmick/:id", setGimmick)
	}
	
	http.Handle("/", router)
}

type Gimmicks struct {
	Gimmicks []Gimmick `json:gimmicks`
}

type Gimmick struct {
  Name string	`json:"name"`
	Description string	`json:"description"`
	Src string	`json:"src"`
	Media string	`json:"media"`
	Env string	`json:"env"`
	Path string	`json:"path"`
	PageAction string	`json:"pageAction"`
	Persona string	`json:"persona`
}

func getGimmicks(ctx *gin.Context){
	gaeCtx = appengine.newContext(ctx.Request)
	
	var gimmicks Gimmicks
	keys, err := datastore.NewQuery("Gimmick").GetAll(gaeCtx, &gimmicks.Gimmicks)
	ctx.JSON(200, json.Marshal(gimmicks));
}

func newGimmick(ctx *gin.Context){
	var gimmick Gimmick
	ctx.BindJson(gimmick)

	gaeCtx = appengine.newContext(gin.Context.Request)
	key := datastore.NewKey(gaeCtx, "Gimmick", "", 0, nil)
	datastore.Put(gaeCtx, key, &gimmick);
}

func getGimmick(ctx *gin.Context){
	id := ctx.params("id");
	var gimmick Gimmick
	gaeCtx = appengine.newContext(gin.Context.Request)
	datastore.Get(gaeCtx, id, &gimmick)

	ctx.JSON(200, json.Marshal(gimmick));
}

func setGimmick(ctx *gin.Context){
	id := ctx.params("id");
	var gimmick Gimmick
	ctx.BindJson(gimmick)

	gaeCtx = appengine.newContext(gin.Context.Request)
	datastore.Put(gaeCtx, id, &gimmick)
}

func test(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "test!",
	})
}