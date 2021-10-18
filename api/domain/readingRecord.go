package domain

type ReadingRecord struct {
	Intersection
	UserId string `json:"userId"`
	Book  Book  `json:"book"`
	PageCount int `json:"pageCount"`
	Dulation int `json:"dulation"`
	RegisteredAt ParsedDateTime `json:"registeredAt"`
}
