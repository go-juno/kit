package kerrors

import (
	"fmt"

	"golang.org/x/xerrors"
)

type ErrBussiness struct {
	msg string
}

func (e *ErrBussiness) Error() string {
	return e.msg
}

func Errorf(format string, a ...interface{}) error {
	return xerrors.Errorf("%w", &ErrBussiness{
		msg: fmt.Sprintf(format, a...),
	})
}
