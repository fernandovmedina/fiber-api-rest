package models

type Book struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Autor     string `json:"autor"`
	Editorial string `json:"editorial"`
}
