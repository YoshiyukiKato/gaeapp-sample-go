package model

import (
	"./base"
)

//Gimmicks is a list of Gimmick.
type Gimmicks struct {
	Items []Gimmick `json:"items"`
}

//Gimmick data model.
type Gimmick struct{
	base.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Src         string `json:"src"`
	Media       string `json:"media"`
	Env         string `json:"env"`
	Path        string `json:"path"`
	PageAction  string `json:"pageAction"`
	Persona     string `json:"persona"`
}

/*
type Gimmick struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Src         string `json:"src"`
	Media       string `json:"media"`
	Env         string `json:"env"`
	Path        string `json:"path"`
	PageAction  string `json:"pageAction"`
	Persona     string `json:"persona"`
}
*/