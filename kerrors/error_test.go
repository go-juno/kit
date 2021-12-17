package kerrors_test

import (
	"fmt"
	"testing"

	kerror "git.yupaopao.com/ops-public/kit/error"
	"git.yupaopao.com/ops-public/kit/kerrors"
)

func TestError(t *testing.T) {

	err := kerrors.Errorf("%s", "测试错误")
	err = kerror.Unwrap(err)
	if e, ok := err.(*kerrors.ErrBussiness); ok {
		fmt.Printf("判断成功")
		fmt.Println(e.Error())
	}

}
