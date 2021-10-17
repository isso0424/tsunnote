package domain

type ReadingRecord struct {
	Intersection
	Book  Book  `json:"book"`
	Quota Quota `json:"quota"`
}
