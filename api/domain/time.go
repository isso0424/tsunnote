package domain

import "time"

var utcLoc = time.FixedZone("UTC", 0)

type ParsedDateTime string

func ParseDateTime(t time.Time) ParsedDateTime {
	return ParsedDateTime(t.In(utcLoc).Format(time.RFC3339))
}

type ParsedDate string

func ParseDate(t time.Time) ParsedDate {
	return ParsedDate(t.In(utcLoc).Format("2006-01-02"))
}

func NowDateTime() ParsedDateTime {
	return ParseDateTime(time.Now().UTC())
}
