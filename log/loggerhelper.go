package log

import (
	"log"
	"strings"

	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

func filterLoggerFlags(args []string, keep bool) []string {
	rlt := make([]string, 0)

	for _, arg := range args {
		if strings.HasPrefix(arg, "--logger.") {
			if keep {
				rlt = append(rlt, arg)
			}
		} else {
			if !keep {
				rlt = append(rlt, arg)
			}
		}
	}

	return rlt
}

type redirectSysLog struct {
	entry *logrus.Entry
}

// CollectSysLog redirects system entry into logrus info,only can call once
func CollectSysLog() {
	log.SetFlags(log.Llongfile)
	// SetLevel and SetFormatter must be called before getActivateHooks.
	logger := &logrus.Logger{
		Out: os.Stdout,
		Formatter: &ClassicFormatter{
			IgnoreFields: []string{"service"},
		},
		Hooks:        nil,
		Level:        qEntry.Logger.Level,
		ExitFunc:     qEntry.Logger.ExitFunc,
		ReportCaller: false,
	}

	logger.SetOutput(ioutil.Discard)
	logger.ReplaceHooks(getActivateHooks())
	entry := logrus.NewEntry(logger).WithField("service", _ServiceName)
	log.SetOutput(&redirectSysLog{entry})
}

func (r *redirectSysLog) Write(p []byte) (n int, err error) {
	r.entry.Println(string(p))
	return len(p), nil
}
