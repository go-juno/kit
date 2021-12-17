package log

import (
	"errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

var _ServiceName string

func New(serviceName string, level Level) (*logrus.Entry, error) {
	if len(serviceName) == 0 {
		return nil, errors.New("service is required")
	}
	parseLevel, err := logrus.ParseLevel(string(level))
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	qEntry.Logger.SetReportCaller(true)
	qEntry.Logger.SetLevel(parseLevel)
	entry := qEntry.WithField("service", serviceName)
	// qEntry 全局增加 service 名称
	qEntry = qEntry.WithField("service", serviceName)
	_ServiceName = serviceName
	return entry, nil
}
