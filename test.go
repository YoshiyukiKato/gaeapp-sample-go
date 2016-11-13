package main

import (
  "fmt"
  "reflect"
)

type Model struct{
  Instance interface{}
}

func (m Model) Create() interface{}{
  schema := reflect.TypeOf(m.Instance)
  instance := reflect.New(schema).Elem()

  for i := 0; i < schema.NumField(); i++ {
    // フィールドの取得
    field := schema.Field(i)
    fmt.Println(field.Name)
    fmt.Println(field.Type)
    fmt.Println(field.Tag)
  }

  return instance
}

type Person struct{
  Name string
  Age int
}

func main(){
  m := Model{ Person{} }
  i := m.Create()
  fmt.Printf("%+v", i)
}
