package err

import "fmt"

func (e Error) Error() string {
	switch e.errorCode {
	case NotFound:
		return fmt.Sprintf("Notfound: %s", e.detail["target"])
	case UnsupportedData:
		return fmt.Sprintf("Unsupported data: %s", e.detail["data"])
	case InvalidValue:
		return fmt.Sprintf("Invalid value: %s", e.detail["value"])
	case Internal:
		return fmt.Sprintf("Internal server error: %s", e.detail["error"])
	default:
		return "Unknown error"
	}
}

func (e Error) ErrorToClient() string {
	if e.errorCode == Internal {
		return "Internal server error"
	}

	return e.Error()
}
