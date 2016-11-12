package model

import (
  "net/http"
  "errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)


//Model is interface for datastore
type Model struct{
  Kind string
  Context appengine.Context
  Schema 
}

func (m Model) Init(r *http.Request){
	m.Context = appengine.NewContext(r)
  return m
}

func (m Model) All() []interface{}, Error {
  if m.Context == nil {
    return nil, errors.New("Context is not initialized")
  }
}

func (m Model) Find(id) interface{}, Error {
  if m.Context == nil {
    return nil, errors.New("Context is not initialized")
  }

  id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

  var instance m.Schema
	if err = datastore.Get(m.Context, key, &instance); err != nil {
		return nil, err
	}

  return instance, err
}

func (m Model) FindBy(paramName, paramValue) interface{}, Error{
  if m.Context == nil {
    return nil, errors.New("Context is not initialized")
  }

  instances, err := datastore.NewQuery(m.Kind).Filter(paramName, paramValue).Run()
	if err != nil {
		return nil, err
	}

  return instances, nil
}

func (m Model) New() interface{}, Error{

}

type Instance struct{}

func (i Instance) Save() int, Error {

}