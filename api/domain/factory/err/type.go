package err

import "fmt"

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

type invalidData struct {
	data string
}

func (e invalidData) Error() string {
	return fmt.Sprintf("Invalid data: ", e.data)
}

func (e invalidData) GetCode() errorCode {
	return InvalidData
}

func (e invalidData) GetDetail() map[string]string {
	return map[string]string{"data": e.data}
}

type constraintViolate struct {
	field string
	constraint string
}

func (e constraintViolate) Error() string {
	return fmt.Sprintf("ConstraintViolate: %s must be %s", e.field, e.constraint)
}

func (e constraintViolate) GetCode() errorCode {
	return ConstraintViolate
}

func (e constraintViolate) GetDetail() map[string]string {
	return map[string]string{"field": e.field, "constraint": e.constraint}
}

type internal struct {
	e error
}

func (e internal) Error() string {
	return e.e.Error()
}

func (e internal) GetCode() errorCode {
	return Internal
}

func (e internal) GetDetail() map[string]string {
	return map[string]string{"error": e.e.Error()}
}
