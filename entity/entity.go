package entity

//Tip holds the properties  of a health tips
type Tip struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Number int64  `json:"number"`
}
