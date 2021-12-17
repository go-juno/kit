package log

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

var _ = func() interface{} {
	registerFormatter("text", reflect.TypeOf(logrus.TextFormatter{}))

	return nil
}()
