package model

//Gimmicks is a list of Gimmick.
var Gimmick Model = Model{ gimmick{} }

type Gimmicks struct {
	Items []gimmick `json:"items"`
}
//Gimmick data model.
type gimmick struct{
	Instance
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