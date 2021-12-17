package kerror

import "golang.org/x/xerrors"

// Unwrap   错误解构
func Unwrap(err error) (uerr error) {
	for err != nil {
		uerr = err
		err = xerrors.Unwrap(err)
	}
	return

}
