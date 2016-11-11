package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	
	"model"
)
//"google.golang.org/appengine/log"

func init() {
	router := gin.New()
	v1 := router.Group("/v1")
	{
		v1.GET("/gimmick/list", getGimmicks)
		v1.POST("/gimmick/new", newGimmick)
		v1.GET("/gimmick/get/:id", getGimmick)
		v1.POST("/gimmick/set/:id", setGimmick)
	}

	http.Handle("/", router)
}

func getGimmicks(ctx *gin.Context) {
	gaeCtx := appengine.NewContext(ctx.Request)
	var gimmicks model.Gimmicks
	
	if _, err := datastore.NewQuery("Gimmick").GetAll(gaeCtx, &gimmicks.Items); err != nil {
		ctx.JSON(500, gin.H{"message": err})
	}
	
	ctx.JSON(200, gimmicks)
}

func newGimmick(ctx *gin.Context) {
	var gimmick model.Gimmick
	if jsonerr := ctx.BindJSON(&gimmick); jsonerr != nil{
		ctx.JSON(400, gin.H{"message": "no data given"})
		return;
	}	

	gaeCtx := appengine.NewContext(ctx.Request)
	newkey := datastore.NewKey(gaeCtx, "Gimmick", "", 0, nil)

	if _, puterr := datastore.Put(gaeCtx, newkey, &gimmick); puterr != nil {
		ctx.JSON(500, gin.H{"message": "put error"})
		return;
	} 
	
	ctx.JSON(200, gin.H{"message": "complete!"})
	
}

func getGimmick(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err})
		return
	}

	var gimmick model.Gimmick
	gaeCtx := appengine.NewContext(ctx.Request)
	key := datastore.NewKey(gaeCtx, "Gimmick", "", id, nil)
	
	if err = datastore.Get(gaeCtx, key, &gimmick); err != nil {
		ctx.JSON(404, gin.H{"message": "no such gimmick"})
		return
	}
	
	ctx.JSON(200, gimmick)
	
}

func setGimmick(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err})
	}

	var gimmick model.Gimmick
	ctx.BindJSON(&gimmick)
	gaeCtx := appengine.NewContext(ctx.Request)
	key := datastore.NewKey(gaeCtx, "Gimmick", "", id, nil)

	if _, err = datastore.Put(gaeCtx, key, &gimmick); err != nil {
		ctx.JSON(500, gin.H{"message": err})
	}

	ctx.JSON(200, gin.H{"message": "complete!"})
}