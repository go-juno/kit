package log

import (
	"os"
	"reflect"
)

const (
	keyStdoutEnabled = "logger.stdout.enabled"
	keyStdoutLevel   = "logger.stdout.level"
)

// StdoutHook output message to StdoutHook
type StdoutHook struct {
	BaseHook
}

// Setup function for StdoutHook
func (h *StdoutHook) Setup() error {

	// stdout 默认 formatter
	h.formatter = &ClassicFormatter{
		TruncateCallerPath: true,
		CallerPathStrip:    true,
		TimestampFormat:    longTimeStamp,
		DisableTimestamp:   false,
		IgnoreFields:       []string{"service"},
	}
	h.baseSetup()

	h.writer = os.Stdout

	return nil
}

var _ = func() interface{} {
	cli.Bool(keyStdoutEnabled, true, "logger.stdout.enabled")
	cli.String(keyStdoutLevel, "", "logger.stdout.level") // DONOT set default level in pflag

	registerHook("stdout", reflect.TypeOf(StdoutHook{}))
	return nil
}()
