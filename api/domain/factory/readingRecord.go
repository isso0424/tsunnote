package factory

import (
	"isso0424/tsunnote_api/domain"
	"isso0424/tsunnote_api/domain/factory/err"
	"time"

	"github.com/google/uuid"
)

type RecordFactory struct {}

func NewRecordFactory() RecordFactory {
	return RecordFactory{}
}

func (f RecordFactory) Generate(
	userId string,
	pageCount,
	dulation int,
	book domain.Book,
) (domain.ReadingRecord, err.Error) {
	id := uuid.New().String()

	return domain.ReadingRecord{
		Intersection: domain.Intersection{
			ID: id,
		},
		PageCount: pageCount,
		Dulation: dulation,
		UserId: userId,
		Book: book,
		RegisteredAt: domain.NowDateTime(),
	}, nil
}

func (f RecordFactory) Reconstruct(
	id,
	userId string,
	pageCount,
	dulation int,
	book domain.Book,
	registeredAt time.Time,
) (domain.ReadingRecord, err.Error) {
	return domain.ReadingRecord{
		Intersection: domain.Intersection{
			ID: id,
		},
		PageCount: pageCount,
		Dulation: dulation,
		UserId: userId,
		Book: book,
		RegisteredAt: domain.ParseDateTime(registeredAt),
	}, nil
}
