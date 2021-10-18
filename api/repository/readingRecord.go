package repository

import (
	"isso0424/tsunnote_api/domain"
	"isso0424/tsunnote_api/repository/err"
)

type RecordRepository interface {
	Create(userId string, book domain.Book, pageCount, dulation int) (domain.ReadingRecord, err.Error)

	FetchByDate(userId string, date domain.ParsedDate) (domain.ReadingRecord, err.Error)
	FetchInWeek(userId string, date domain.ParsedDate) (domain.ReadingRecord, err.Error)

	Delete(userId string) (domain.ReadingRecord, err.Error)
}
