package entity

//Post holds the definition of a Post item
type Post struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Number int64  `json:"number"`
}
