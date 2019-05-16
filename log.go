package logger

import (
	"fmt"
)

type Log struct {
	hasPrefix bool
	prefixStr string
}

func NewLog() *Log {
	return &Log{
		hasPrefix: true,
		prefixStr: GetPrefix(),
	}
}

func (this *Log) getPrefix() string {
	if this.hasPrefix == false {
		this.prefixStr = GetPrefix()
		this.hasPrefix = true
	}

	return this.prefixStr
}

func (this *Log) Debug(v ...interface{}) {
	Output(3, "DEBUG", this.getPrefix(), v...)
}

func (this *Log) Debugf(format string, v ...interface{}) {
	Output(3, "DEBUG", this.getPrefix(), fmt.Sprintf(format, v...))
}

func (this *Log) Info(v ...interface{}) {
	Output(3, "INFO", this.getPrefix(), v...)
}

func (this *Log) Infof(format string, v ...interface{}) {
	Output(3, "INFO", this.getPrefix(), fmt.Sprintf(format, v...))
}

func (this *Log) Warn(v ...interface{}) {
	Output(3, "WARN", this.getPrefix(), v...)
}

func (this *Log) Warnf(format string, v ...interface{}) {
	Output(3, "WARN", this.getPrefix(), fmt.Sprintf(format, v...))
}

func (this *Log) Error(v ...interface{}) {
	Output(3, "ERROR", this.getPrefix(), v...)
}

func (this *Log) Errorf(format string, v ...interface{}) {
	Output(3, "ERROR", this.getPrefix(), fmt.Sprintf(format, v...))
}

func (this *Log) Output(calldepth int, s string) error {
	Output(2+calldepth, "WARN", this.getPrefix(), s)
	return nil
}