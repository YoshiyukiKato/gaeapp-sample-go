package model

import (
  "reflect"
  "errors"
  "golang.org/x/net/context"
)

/*
  "net/http"
  "google.golang.org/appengine"
  "google.golang.org/appengine/datastore"
*/


//Model
type Model struct{
  Kind string
  Context *context.Context
  Columns Columns
}

func (m Model) SetParams(params Params, i Instance){
  for colName, column := range m.Columns{
    if params[colName] != nil {
      if param, err := column.Validate(params[colName]); err == nil {
        i.Params[colName] = param
      }
    }
  }
}

func (m Model) New(params Params) (i Instance){
  i = Instance{ m, Params{} }
  m.SetParams(params, i)
  return i
}

//Column
type Columns map[string]Column
type Column struct{
  example interface{}
  defaultTo interface{}
  notNull bool
  constraint (func(interface{}) bool)
}

func (c Column) Example(value interface{}) Column{
  c.example = value
  return c
}

func (c Column) DefaultTo(value interface{}) Column{
  c.defaultTo = value
  return c
}

func (c Column) NotNull(value bool) Column{
  c.notNull = value
  return c
}

func (c Column) Constraint(value (func(interface{})bool)) Column{
  c.constraint = value
  return c
}

func (c Column) Validate(value interface{}) (validated interface{}, err error){
  //1. validate nil 
  if value == nil && c.notNull{
    if c.defaultTo == nil{
      err = errors.New("unexpected nil")
    }else{
      value = c.defaultTo
    }
  }
  //2. validate type.
  if reflect.TypeOf(c.example) != reflect.TypeOf(value){
    err = errors.New("unexpected type")
    return nil, err
  }
  //3. validate by Constraint.
  if c.constraint != nil{
    if c.constraint != nil && !c.constraint(value){
      err = errors.New("unexpected value")
      return nil, err
    }
  }

  validated = value
  return validated, nil  
}

//Instance
type Instance struct{
  Model Model
  Params Params
}

func (i Instance) SetParams(params Params){
  i.Model.SetParams(params, i)
}

//Params
type Params map[string]interface{}

/*
func (m Model) Init(r *http.Request) Model{
  m.Context = appengine.NewContext(r)
  return m
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
*/