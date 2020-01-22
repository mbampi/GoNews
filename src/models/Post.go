package models

// Post struct
type Post struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Author  *Author `json:"author"`
}

func getAll() {
	//con.Exec("INSERT INTO post VALUES(2, 'Post from GO', 'Post made from GOlang');")

}
