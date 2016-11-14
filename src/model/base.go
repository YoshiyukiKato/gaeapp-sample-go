package model

import (
  "net/http"
  "reflect"
  "errors"
  "golang.org/x/net/context"
  "google.golang.org/appengine"
  "google.golang.org/appengine/datastore"
)

//too many error. stash
//Model is interface for datastore
type Model struct{
  Kind string
  Context context.Context
  Instance interface{}
}

type Instance struct{
  Key datastore.Key
}

func (m Model) Init(r *http.Request) Model{
  m.Context = appengine.NewContext(r)
  return m
}

func (m Model) New(params map[string]interface{}) (instance Instance, err error){
  schema := reflect.TypeOf(m.Instance)
  instance = reflect.New(schema).Elem()
  instance.Key = datastore.NewKey(m.Context, m.Kind, "", 0, nil)

  for i := 0; i < schema.NumField(); i++ {
    field := schema.Field(i)
    param := params[field.Name]
    if param != nil && reflect.TypeOf(param) == field.Type{
      instance.Field(i).Set(param)
    }
  }

  return instance, nil
}

func (m Model) Save(instance Instance) (err error){
  if m.Context == nil {
    return errors.New("Context is not initialized")
  }

  instanceKey := instance.Key
  if instanceKey == nil {
  }

  key, err := datastore.Put(m.Context, instanceKey, &instance)
  if key != nil {
    instance.Key = key
  }

  return err
}

func (m Model) Find(key datastore.Key) (instance reflect.Value, err error){
  if m.Context == nil {
    return nil, errors.New("Context is not initialized")
  }
  
  schema := reflect.TypeOf(m.Instance)
  instance = reflect.New(schema).Elem()
  if geterr := datastore.Get(m.Context, key, instance); geterr == nil{
    instance.Key = key
  }
  
  return instance, geterr
}

func (m Model) All() (instances []interface{}, err error){
  return m.Where(nil)
}

func (m Model) FindBy(paramName string, paramValue interface{}) (instances []reflect.Value, err error){
  terms := map[string]interface{}{
    paramName + " =" : paramValue,
  }

  return m.Where(terms)  
}

func (m Model) Where(terms map[string]interface{}) (instances []reflect.Value, err error){
  if m.Context == nil {
    return nil, errors.New("Context is not initialized")
  }

  query := datastore.NewQuery(m.Kind)
  if terms != nil{
    for filterStr, value := range terms{
      query.Filter(filterStr, value)
    }
  }
  
  var keys []datastore.Key //m.Shemaをどう持たせるか問題
  _, err = query.GetAll(keys)
  if keyerr != nil{
    return nil, keyerr  
  }

  instanceType := reflect.TypeOf(m.Instance)
  instances := reflect.SliceOf(instanceType)
  geterr = datastore.GetMulti(m.Context, keys, instances)
  return instances, geterr
}