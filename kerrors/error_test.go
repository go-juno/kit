package kerrors_test

import (
	"fmt"
	"testing"

	kerror "github.com/go-juno/kit/error"
	"github.com/go-juno/kit/kerrors"
)

func TestError(t *testing.T) {

	err := kerrors.Errorf("%s", "测试错误")
	err = kerror.Unwrap(err)
	if e, ok := err.(*kerrors.ErrBussiness); ok {
		fmt.Printf("判断成功")
		fmt.Println(e.Error())
	}

}
