package domain

type Book struct {
	Intersection
	ISBN  string `json:"isbn"`
	Title string `json:"title"`
	PageCount int `json:"pageCount"`
}
