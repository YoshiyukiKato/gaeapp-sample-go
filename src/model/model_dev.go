package model

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)


type Instance struct{}

//Model is interface for datastore
type Model struct{
  Kind string
}

func (m Model) all() []Instance{

}

func (m Model) all() []Instance{

}


func (m Model) find(id) Instance {

}

func (m Model) findBy(paramName, id) Instance {

}