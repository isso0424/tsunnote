package err

import "fmt"

func (e Error) Error() string {
	switch e.Code {
	case NotFound:
		return fmt.Sprintf("Notfound: %s", e.Detail["target"])
	case UnsupportedData:
		return fmt.Sprintf("Unsupported data: %s", e.Detail["data"])
	case InvalidValue:
		return fmt.Sprintf("Invalid value: %s", e.Detail["value"])
	case Internal:
		return fmt.Sprintf("Internal server error: %s", e.Detail["error"])
	default:
		return "Unknown error"
	}
}

func (e Error) ErrorToClient() string {
	if e.Code == Internal {
		return "Internal server error"
	}

	return e.Error()
}
