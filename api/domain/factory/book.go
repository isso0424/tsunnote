package factory

import (
	"fmt"
	"isso0424/tsunnote_api/domain"
	"isso0424/tsunnote_api/domain/factory/err"

	"github.com/google/uuid"
)

const MAX_TITLE_LEN = 64

type bookAPIFactory interface {
	GenerateWithISBN(isbnCode string, pageCount int) (domain.Book, err.Error)
}

type BookFactory struct {
	apiFactory bookAPIFactory
}

func NewBookFactory(f bookAPIFactory) BookFactory {
	return BookFactory{ apiFactory: f }
}

func (f BookFactory) Generate(title string, pageCount int) (domain.Book, err.Error) {
	id := uuid.New().String()

	tLen := len([]rune(title))
	if tLen > MAX_TITLE_LEN {
		return domain.Book{}, err.NewConstraintViolate(fmt.Sprintf("shorter than %d characters", MAX_TITLE_LEN), "title")
	}

	return domain.Book{
		Intersection: domain.Intersection{
			ID: id,
		},
		Title: title,
		PageCount: pageCount,
		ISBN: "",
		RegisteredAt: domain.NowDateTime(),
	}, nil
}

func (f BookFactory) Reconstruction(id, title, isbn string, pageCount int, registeredAt domain.ParsedDateTime) (domain.Book, err.Error) {
	tLen := len([]rune(title))
	if tLen > MAX_TITLE_LEN {
		return domain.Book{}, err.NewConstraintViolate(fmt.Sprintf("shorter than %d characters", MAX_TITLE_LEN), "title")
	}

	return domain.Book {
		Intersection: domain.Intersection{
			ID: id,
		},
		Title: title,
		PageCount: pageCount,
		ISBN: isbn,
		RegisteredAt: registeredAt,
	}, nil
}
