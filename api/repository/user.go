package repository

import (
	"isso0424/tsunnote_api/domain"
	"isso0424/tsunnote_api/repository/err"
)

type UserRepository interface {
	Create(name, email, passHashed string) (domain.User, err.Error)
	AddNewQuota(userId string, newQuota domain.Quota) (domain.User, err.Error)
	UpdateProfile(name, email string) (domain.User, err.Error)
	Delete(userId string) (domain.User, err.Error)
}
