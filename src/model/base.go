package base

import (
  "net/http"
  "errors"
  "google.golang.org/appengine"
  "google.golang.org/appengine/datastore"
)

//Model is interface for datastore
type Model struct{
  Kind string
  Context appengine.Context
  Instance interface{}
}

type Instance struct{
  Key datastore.Key
}

func (m Model) Init(r *http.Request) Model{
  m.Context = appengine.NewContext(r)
  return m
}

func (m Model) New(params map[string] interface{}) Instance, Error{
  schema := reflect.TypeOf(m.Instance)
  instance := reflect.New(schema).Elem()

  var param interface{}
  for i := 0; i < schema.NumField(); i++ {
    // フィールドの取得
    field := schema.Field(i)
    param = params[field.Name]
    if param != nil && reflect.TypeOf(param) == field.Type{
      instance.Field(i).Set(param)
    }
  }

  return instance
}

func (m Model) Save(instance Instance) Instance, Error{
  if m.Context == nil {
    return nil, errors.New("Context is not initialized")
  }

  if instance.Key == nil {
    newkey := datastore.NewKey(m.Context, m.Kind, "", 0, nil)
  }

  key, err := datastore.Put(m.Context, newkey, &instance)
  if key != nil {
    instance.Key = key
  }

  return instance, err
}

func (m Model) Find(key datastore.Key) interface{}, Error{
  if m.Context == nil {
    return nil, errors.New("Context is not initialized")
  }
  
  schema := reflect.TypeOf(m.Instance)
  instance := reflect.New(schema).Elem()
  if geterr := datastore.Get(m.Context, key, instance); geterr == nil{
    instance.Key = key
  }
  
  return instance, geterr
}

func (m Model) All() []Instance, Error {
  return m.Where(nil)
}

func (m Model) FindBy(paramName string, paramValue interface{}) interface{}, Error{ 
  terms := map["string"]interface{}{
    paramName + " =" : paramValue
  }

  return m.Where(terms)  
}

func (m Model) Where(terms map[string]interface{}) []interface{}, Error{
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
  _, err := query.GetAll(keys)
  if keyerr != nil{
    return nil, keyerr  
  }

  instanceType := reflect.TypeOf(m.Instance)
  instances := reflect.SliceOf(instanceType)
  geterr = datastore.GetMulti(m.Context, keys, instances)
  return instances, geterr
}