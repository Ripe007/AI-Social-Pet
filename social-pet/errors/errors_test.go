package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrors(t *testing.T) {
	fmt.Println(errors.New("err xxx"))

	wrapErr := SystemError.Wrap(errors.New("err xxx"), "这个操作有误，因为")
	fmt.Println(wrapErr, "错误码：", GetType(wrapErr))
}
