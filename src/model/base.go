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
  Columns map[string]Column
}

type Column struct{
  Type string
  NotNull bool
  DefaultTo interface{}
  Constraint func(interface{}) bool
}

func (c Column) validate(value interface{}) (value interface{}, err error){
  //1. validate type.
  //2. validate by Constraint.
  //3. if nil, set default value.
  //4. validate if nil when notNull is true
}

type Instance struct{
  Model Model
  Params map[string]Param
}

type Param{
  Column Column
  Value interface{}
}

func (i Instance) setParams(params map[string]interface{}){
  var param Param
  for colName, column := range i.Model.Columns{
    if params[colName] != nil {
      //TODO validate type of param by colValue
      i.Params[colName] = Param{ column, params[colName] }
    }
  }
}

func (m Model) New(){

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