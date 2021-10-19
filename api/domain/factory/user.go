package factory

import (
	"encoding/json"
	"fmt"
	"isso0424/tsunnote_api/domain"
	"isso0424/tsunnote_api/domain/factory/err"
	"net/mail"
	"time"

	"github.com/google/uuid"
)

type UserFactory struct{}

const MAX_NAME_LENGTH = 32

func NewUserFactory() UserFactory {
	return UserFactory{}
}

type InitialQuota struct {
	SecondsPerDay int32
	TargetWeekDay []int                 `json:"targetWeekDay"`
	CreatedAt     domain.ParsedDateTime `json:"createdAt"`
}

func (f UserFactory) Generate(
	name,
	passHashed,
	email string,
	firstQuota InitialQuota,
) (domain.User, err.Error) {
	_, e := mail.ParseAddress(email)
	if e != nil {
		return domain.User{}, err.NewConstraintViolate("compliant email format", "email")
	}

	nLength := len([]rune(name))
	if nLength > MAX_NAME_LENGTH {
		return domain.User{}, err.NewConstraintViolate(fmt.Sprintf("shorter than %d characters", MAX_NAME_LENGTH), "name")
	}

	id := uuid.New().String()

	shelfId := uuid.New().String()
	shelf := domain.BookShelf{
		Intersection: domain.Intersection{
			ID: shelfId,
		},
		Books: []domain.Book{},
	}

	quotaId := uuid.New().String()
	quota := domain.Quota{
		Intersection: domain.Intersection{
			ID: quotaId,
		},
		SecondsPerDay: firstQuota.SecondsPerDay,
		TargetWeekDay: firstQuota.TargetWeekDay,
		NextQuotaID:   "",
		CreatedAt:     domain.NowDateTime(),
	}

	return domain.User{
		Intersection: domain.Intersection{
			ID: id,
		},
		Name:         name,
		PassHashed:   passHashed,
		Email:        email,
		FirstQuotaID: quota.ID,
		Quotas:       []domain.Quota{quota},
		BookShelf:    shelf,
	}, nil
}

type RawQuota struct {
	ID            string
	SecondsPerDay int32
	TargetWeekDay string
	NextQuotaID   string
	CreatedAt     time.Time
}

func parseQuotaList(firstQuotaId string, quotas []RawQuota) ([]domain.Quota, err.Error) {
	nextID := firstQuotaId
	result := []domain.Quota{}
	for nextID != "" {
		for _, quota := range quotas {
			if quota.ID == nextID {
				tWeekD := []int{}

				e := json.Unmarshal([]byte(quota.TargetWeekDay), &tWeekD)
				if e != nil {
					return []domain.Quota{}, err.NewInternal(e)
				}

				result = append(
					result,
					domain.Quota{
						Intersection: domain.Intersection{
							ID: quota.ID,
						},
						SecondsPerDay: quota.SecondsPerDay,
						TargetWeekDay: tWeekD,
						NextQuotaID:   quota.NextQuotaID,
						CreatedAt:     domain.ParseDateTime(quota.CreatedAt),
					},
				)

				nextID = quota.NextQuotaID
				break
			}
		}
	}

	return result, nil
}

func (f UserFactory) Reconstruction(
	id,
	name,
	passHashed,
	email,
	firstQuotaId,
	bookShelfId string,
	books []domain.Book,
	quotas []RawQuota,
) (domain.User, err.Error) {
	quotaList, e := parseQuotaList(firstQuotaId, quotas)
	if e != nil {
		return domain.User{}, e
	}

	shelf := domain.BookShelf{
		Intersection: domain.Intersection{
			ID: bookShelfId,
		},
		Books: books,
	}

	return domain.User{
		Intersection: domain.Intersection{
			ID: id,
		},
		Name:         name,
		PassHashed:   passHashed,
		Email:        email,
		FirstQuotaID: firstQuotaId,
		BookShelf:    shelf,
		Quotas:       quotaList,
	}, nil
}
