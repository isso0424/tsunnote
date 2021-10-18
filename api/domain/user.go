package domain

type User struct {
	Intersection
	Name       string  `json:"name"`
	PassHashed string  `json:"passHashed"`
	Email      string  `json:"email"`
	FirstQuotaID string `json:"firstQuotaId"`
	Quotas     []Quota `json:"quotas"`
}
