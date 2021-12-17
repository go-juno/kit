package log

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

const (
	longTimeStamp = "2006/01/02 15:04:05.000000Z07:00"
)

var (
	gRegisteredFormatters = make(map[string]reflect.Type)
)

func registerFormatter(name string, typ reflect.Type) {
	gRegisteredFormatters[name] = typ
}

func newFormatter(name string, key string) (logrus.Formatter, error) {
	var err error
	var typ reflect.Type
	var ok bool

	if typ, ok = gRegisteredFormatters[name]; !ok {
		return nil, fmt.Errorf("[log] formatter name(%s) not registered", name)
	}

	f := reflect.New(typ)

	if err = v.UnmarshalKey(key, f.Interface()); err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	return f.Interface().(logrus.Formatter), nil
}
