package err

func NewNotFound(target string) Error {
	return Error{
		errorCode: NotFound,
		detail: map[string]string{ "target": target },
	}
}

func NewUnsupportedData(data string) Error {
	return Error{
		errorCode: UnsupportedData,
		detail: map[string]string{ "data": data },
	}
}

func NewInvalidValue(value string) Error {
	return Error{
		errorCode: InvalidValue,
		detail: map[string]string{ "value": value },
	}
}

func NewInternal(e error) Error {
	return Error{
		errorCode: Internal,
		detail: map[string]string{ "error": e.Error() },
	}
}
