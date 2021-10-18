package domain

type Quota struct {
	Intersection
	SecondsPerDay int32  `json:"targetSecondsPerDay"`
	TargetWeekDay []int  `json:"targetWeekDay"`
	NextQuotaID   string `json:"nextQuotaID"`
	CreatedAt     ParsedDateTime `json:"createdAt"`
}
