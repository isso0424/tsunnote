package repository

import (
	"isso0424/tsunnote_api/domain"
	"isso0424/tsunnote_api/repository/err"
)

type UserRepository interface {
	Create(name, email, passHashed string) (domain.User, err.Error)

	AddNewQuota(userId string, newQuota domain.Quota) (domain.User, err.Error)
	AddNewBook(userId string, newBook domain.Book) (domain.User, err.Error)
	UpdateProfile(userId, name, email string) (domain.User, err.Error)

	FetchByID(id string) (domain.User, err.Error)

	Delete(userId string) (domain.User, err.Error)
}
