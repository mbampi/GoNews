package models

// Post struct
type Post struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Author  *Author `json:"author"`
}

// Author struct (cpf as primary key)
type Author struct {
	CPF  int    `json:"cpf"`
	Name string `json:"name"`
}
