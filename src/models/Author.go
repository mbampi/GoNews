package models

// Author struct (cpf as primary key)
type Author struct {
	CPF  int    `json:"cpf"`
	Name string `json:"name"`
}
