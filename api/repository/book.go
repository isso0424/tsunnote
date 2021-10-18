package repository

import (
	"isso0424/tsunnote_api/domain"
	"isso0424/tsunnote_api/repository/err"
)

type bookSortType int

const (
	CreatedAt = iota
	PageCount
)

type BookRepository interface {
	CreateWithISBN(isbnCode string, pageCount int) (domain.Book, err.Error)
	Create(title string, pageCount int) (domain.Book, err.Error)
	Delete(id string) (domain.Book, err.Error)
	FetchByID(id string, sort bookSortType) (domain.Book, err.Error)
	FetchISBNCode(code string, sort bookSortType) (domain.Book, err.Error)
}
