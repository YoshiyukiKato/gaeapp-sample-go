package main

import (
  "fmt"
  "reflect"
  "errors"
)

type Model struct{
  Kind string
  Schema interface{} //雛形
}

func (m Model) New(params map[string]interface{}) (instance Instance, err error){
  schema := reflect.New(reflect.TypeOf(m.Schema)).Elem()
  instance = Instance{ m, schema }
  err = instance.SetParams(params)
  return instance, err
}

type Instance struct{
  Model Model
  Params interface{} //実際の値
}

func (i Instance) SetParams(params map[string]interface{}) error{
  for k, v := range params {
    err := i.SetParam(k, v)
    if err != nil {
      return err
    }
  }
  return nil
}

func (i Instance) SetParam(name string, value interface{}) error {
  structValue := reflect.ValueOf(i.Params)
  structFieldValue := structValue.FieldByName(name)//つまりParamsをinterface型にしておくと、FieldByNameとかが使えない
  fmt.Printf("%+v\n", structFieldValue)

  if !structFieldValue.IsValid() {
    return fmt.Errorf("No such field: %s in param", name)
  }

  if !structFieldValue.CanSet() {
    return fmt.Errorf("Cannot set %s field value", name)
  }

  structFieldType := structFieldValue.Type()
  val := reflect.ValueOf(value)
  if structFieldType != val.Type() {
    return errors.New("Provided value type didn't match param field type")
  }

  structFieldValue.Set(val)
  return nil
}

type UserSchema struct{
  Name string
  Age int
}

var User Model = Model{ "user", UserSchema{} }

func main() {
  params := map[string]interface{}{ "Name" : "taro", "Age" : 18 }
  user, err := User.New(params)
  fmt.Printf("%+v\n", user)
  fmt.Printf("%+v\n", err)
}