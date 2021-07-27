package error

import (
	"errors"
)

const (
	InvalidJSONParamMessage = "Invalid Json param"
)

func InvalidJsonParamResponse() error {
	return errors.New(InvalidJSONParamMessage)
}
