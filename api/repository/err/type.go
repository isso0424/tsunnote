package err

type code int

const (
	NotFound = iota
	UnsupportedData
	InvalidValue
	Internal
)

type Error struct {
	error
	errorCode code
	detail map[string]string
}
