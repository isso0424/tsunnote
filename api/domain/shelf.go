package domain

type BookShelf struct {
	Intersection
	Books []Book `json:"books"`
	User User `json:"user"`
}
