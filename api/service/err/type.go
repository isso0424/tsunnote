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
	code errorCode
	detail map[string]string
}
