package err

func NewInvalidData(data string) Error {
	return &invalidData{ data }
}

func NewConstraintViolate(constraint string, field string) Error {
	return &constraintViolate{ field, constraint }
}

func NewInternal(err error) Error {
	return &internal{ e: err }
}
