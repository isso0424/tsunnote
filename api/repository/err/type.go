package err

type errorCode int

const (
	NotFound = iota
	UnsupportedData
	InvalidValue
	Internal
)

type Error struct {
	error
	Code errorCode
	Detail map[string]string
}
