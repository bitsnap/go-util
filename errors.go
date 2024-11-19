package goutil

import (
	"errors"
	"strings"
)

type Errors []error

func (e Errors) String() string {
	sb := strings.Builder{}
	for _, err := range e {
		sb.WriteString(err.Error())
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (e Errors) Error() error {
	return errors.New(e.String())
}
