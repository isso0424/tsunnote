package repository

import (
	"isso0424/tsunnote_api/domain"
	"isso0424/tsunnote_api/repository/err"
)

type BookRepository interface {
	Create(isbnCode, title string, pageCount int) (domain.Book, err.Error)
}
