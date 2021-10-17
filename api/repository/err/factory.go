package err

func NewNotFound(target string) Error {
	return Error{
		Code: NotFound,
		Detail: map[string]string{ "target": target },
	}
}

func NewUnsupportedData(data string) Error {
	return Error{
		Code: UnsupportedData,
		Detail: map[string]string{ "data": data },
	}
}

func NewInvalidValue(value string) Error {
	return Error{
		Code: InvalidValue,
		Detail: map[string]string{ "value": value },
	}
}

func NewInternal(e error) Error {
	return Error{
		Code: Internal,
		Detail: map[string]string{ "error": e.Error() },
	}
}
