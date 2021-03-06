package log

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

// NullFormatter format message to empty string
type NullFormatter struct {
}

// Format function of NullFormatter
func (NullFormatter) Format(e *logrus.Entry) ([]byte, error) {
	return []byte{}, nil
}

var _ = func() interface{} {
	registerFormatter("null", reflect.TypeOf(NullFormatter{}))
	return nil
}()
