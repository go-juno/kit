package log

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

func getLogLevels(baseLevel logrus.Level) (level []logrus.Level) {
	level = make([]logrus.Level, 0)
	for i := baseLevel; i > logrus.PanicLevel; i-- {
		level = append(level, i)
	}
	return
}

// HookSetuper is the base interface a log hook must implement
type HookSetuper interface {
	Setup() error
}

// BaseHook for some common function for hooks in log
type BaseHook struct {
	Name  string
	Level string

	formatter logrus.Formatter
	logLevels []logrus.Level
	writer    io.Writer
}

// Fire output message to hook writer
func (h *BaseHook) Fire(e *logrus.Entry) error {
	// fmt.Println("fire:", h.Name)
	dataBytes, err := h.formatter.Format(e)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	_, err = h.writer.Write(dataBytes)

	return err
}

// Levels return all available debug level of a hook
func (h *BaseHook) Levels() []logrus.Level {
	return h.logLevels
}

func (h *BaseHook) baseSetup() {
	// setup levels
	var level = qEntry.Logger.Level
	var err error
	if h.Level = v.GetString(strings.Join([]string{"logger", h.Name, "level"}, ".")); h.Level != "" {
		if level, err = logrus.ParseLevel(h.Level); err != nil {
			err = xerrors.Errorf("%w", err)
			fmt.Printf("[log] setup hook(%s), parse level fail:%s\n", h.Name, err)
			level = qEntry.Logger.Level
		}
	}

	h.logLevels = getLogLevels(level)
	// setup formatters
	if hookFormatterName := v.GetString(strings.Join([]string{"logger", h.Name, "formatter", "name"}, ".")); hookFormatterName != "" {
		if h.formatter, err = newFormatter(hookFormatterName, strings.Join([]string{"logger", h.Name, "formatter", "opts"}, ".")); err != nil {
			err = xerrors.Errorf("%w", err)
			fmt.Printf("[log] setup hook(%s) formatter(%s) fail:%s\n", h.Name, hookFormatterName, err)
			h.formatter = qEntry.Logger.Formatter
		}
	} else if h.formatter == nil {
		h.formatter = qEntry.Logger.Formatter
	}
}

var gRegisteredHooks = make(map[string]reflect.Type)
var setupHooks = make(map[string]logrus.Hook)

func registerHook(name string, typ reflect.Type) {
	gRegisteredHooks[name] = typ

	if _, ok := reflect.New(typ).Interface().(HookSetuper); !ok {
		panic(fmt.Sprintf("[log] registe hook (%s) fail: must be HookSetuper()", name))
	}
}

func newHook(name string) (logrus.Hook, error) {
	var err error
	var typ reflect.Type
	var ok bool

	if typ, ok = gRegisteredHooks[name]; !ok {
		return nil, fmt.Errorf("[log] hook name(%s) not registered", name)
	}

	h := reflect.New(typ)

	h.Elem().FieldByName("Name").SetString(name)

	hook := h.Interface().(logrus.Hook)

	setuper, _ := hook.(HookSetuper)
	if err = setuper.Setup(); err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	return hook, nil
}
