package main

import (
  "fmt"
  "encoding/json"
)

type parent struct{
  Name string `json:"name"`
  Children []child `json:"children"`
}

type child struct{
  Name string `json:"name"`
}

func main(){
  c := child{ "nick" }
  p := parent{ "jack", []child{c} }
  fmt.Printf("%+v\n", p)
  b, err := json.Marshal(p)
  if err != nil{
    fmt.Print(err)
  }
  fmt.Print(string(b))
}