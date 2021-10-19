package err

type errorCode int

const (
	InvalidData = iota
	ConstraintViolate
	Internal
)

type Error interface {
	error
	GetCode() errorCode
	GetDetail() map[string]string
}
