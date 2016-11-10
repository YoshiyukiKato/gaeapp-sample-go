package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
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

//Gimmick data model.
type gimmick struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Src         string `json:"src"`
	Media       string `json:"media"`
	Env         string `json:"env"`
	Path        string `json:"path"`
	PageAction  string `json:"pageAction"`
	Persona     string `json:"persona"`
}

//Gimmicks is a list of Gimmick.
type gimmicks struct {
	List []gimmick `json:"list"`
}

func getGimmicks(ctx *gin.Context) {
	gaeCtx := appengine.NewContext(ctx.Request)
	var gimmicks gimmicks
	keys, err := datastore.NewQuery("Gimmick").GetAll(gaeCtx, &gimmicks.List)

	if err != nil {
		ctx.JSON(500, gin.H{"message": err})
	} else {
		ctx.JSON(200, json.Marshal(gimmicks))
	}
}

func newGimmick(ctx *gin.Context) {
	var gimmick gimmick
	ctx.BindJSON(&gimmick)
	gaeCtx := appengine.NewContext(ctx.Request)
	newkey := datastore.NewKey(gaeCtx, "Gimmick", "", 0, nil)
	key, err := datastore.Put(gaeCtx, newkey, &gimmick)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err})
	} else {
		ctx.JSON(200, gin.H{"message": "complete!"})
	}
}

func getGimmick(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err})
	}

	var gimmick gimmick
	gaeCtx := appengine.NewContext(ctx.Request)
	key := datastore.NewKey(gaeCtx, "Gimmick", "", id, nil)
	_key, err := datastore.Get(gaeCtx, key, &gimmick)

	if err != nil {
		ctx.JSON(500, gin.H{"message": err})
	} else {
		ctx.JSON(200, json.Marshal(gimmick))
	}
}

func setGimmick(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err})
	}

	var gimmick gimmick
	ctx.BindJSON(&gimmick)
	gaeCtx := appengine.NewContext(ctx.Request)
	key := datastore.NewKey(gaeCtx, "Gimmick", "", id, nil)
	_key, err := datastore.Put(gaeCtx, key, &gimmick)

	if err != nil {
		ctx.JSON(500, gin.H{"message": err})
	} else {
		ctx.JSON(200, gin.H{"message": "complete!"})
	}
}
