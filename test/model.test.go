package main

import (
  "fmt"
  "model"
)
//"golang.org/x/net/context"

func initUser() model.Model{
  var User model.Model = model.Model{}
  User.Kind = "user" 
  User.Columns = model.Columns{
    "name" : model.Column{}.Example("name").NotNull(true),
    "age" : model.Column{}.Example(0).NotNull(false),
  }
  return User
}

func main(){
  User := initUser()
  params := map[string]interface{}{
    "name" : "taro",
    "age" : 18,
  }
  user := User.New(params)
  fmt.Printf("%+v\n", user)
  user.SetParams(model.Params{ "age" : 20 })
  fmt.Printf("%+v\n", user)
}