package err

import repoErr "isso0424/tsunnote_api/repository/err"

func NewNotFound(target string) Error {
	return Error{
		Code: NotFound,
		Detail: map[string]string{ "target": target },
	}
}

func NewInvalidData(arg string) Error {
	return Error{
		Code: InvalidData,
		Detail: map[string]string{ "arg": arg },
	}
}

func NewForbidden(process string) Error {
	return Error{
		Code: Forbidden,
		Detail: map[string]string{ "process": process },
	}
}

func NewFromRepositoryError(e repoErr.Error) Error {
	switch e.Code {
	case repoErr.Internal:
		return Error{
			Code: Internal,
			Detail: e.Detail,
		}
	case repoErr.InvalidValue, repoErr.UnsupportedData:
		return Error{
			Code: InvalidData,
			Detail: e.Detail,
		}
	case repoErr.NotFound:
		return Error{
			Code: NotFound,
			Detail: e.Detail,
		}
	default:
		return Error{
			Code: Internal,
			Detail: e.Detail,
		}
	}
}
