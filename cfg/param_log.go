package cfg

import (
	"github.com/ameteiko/golang-kit/errors"
	"github.com/ameteiko/golang-kit/log"
)

//
// LogInfoProvider declares all the log info getters.
//
type LogInfoProvider interface {
	ParameterInfoProvider

	GetHost() string
}

//
// LogInfo is a log config parameter.
//
type LogInfo struct {
	severity string

	*StringParameter
}

//
// validate validates the URL config parameter.
//
func (l *LogInfo) validate() error {
	var err error
	if err = l.StringParameter.validate(); nil != err {
		return err
	}
	severity := l.GetValue()
	if severity != log.SeverityDebug && severity != log.SeverityError && severity != log.SeverityInfo {
		return errors.WithMessage(ErrLogSeverityIncorrectValue, `kit-cfg@LogInfo.validate [value (%s)]`, severity)
	}

	l.severity = severity

	return nil
}

//
// GetSeverity returns log severity setting.
//
func (l *LogInfo) GetSeverity() string {

	return l.severity
}
