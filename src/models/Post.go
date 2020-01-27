package models

// Post struct
type Post struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Author  *Author `json:"author"`
}
