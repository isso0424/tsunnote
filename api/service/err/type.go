package err

type errorCode int

const (
	NotFound = iota
	InvalidData
	Forbidden
	Internal
)

type Error struct {
	error
	Code   errorCode
	Detail map[string]string
}
