package domain

type User struct {
	Intersection
	Name       string `json:"name"`
	PassHashed string `json:"passHashed"`
	Email      string `json:"email"`
}
