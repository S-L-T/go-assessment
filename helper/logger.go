package helper

import (
	"github.com/sirupsen/logrus"
	"runtime/debug"
	"strconv"
)

type LogLevel uint8

const (
	PanicLevel LogLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func (l LogLevel) String() string {
	n := []string{"panic", "fatal", "error", "warn", "info", "debug", "trace"}
	i := uint8(l)
	switch {
	case i <= uint8(TraceLevel):
		return n[i]
	case i > 0:
		return n[0]
	default:
		return strconv.Itoa(int(i))
	}
}

func InitializeLogger(minimumLevel LogLevel) error {
	logrus.SetReportCaller(true)
	l, err := logrus.ParseLevel(minimumLevel.String())
	if err != nil {
		return err
	}

	logrus.SetLevel(l)

	return nil
}

func Log(err error, level LogLevel) {
	switch level {
	case PanicLevel:
		logrus.WithField("stack", string(debug.Stack())).Panic(err)
		break
	case FatalLevel:
		logrus.WithField("stack", string(debug.Stack())).Fatal(err)
		break
	case ErrorLevel:
		logrus.WithField("stack", string(debug.Stack())).Error(err)
		break
	case WarnLevel:
		logrus.WithField("stack", string(debug.Stack())).Warn(err)
		break
	case InfoLevel:
		logrus.WithField("stack", string(debug.Stack())).Info(err)
		break
	case DebugLevel:
		logrus.WithField("stack", string(debug.Stack())).Debug(err)
		break
	case TraceLevel:
		logrus.WithField("stack", string(debug.Stack())).Trace(err)
		break
	default:
		logrus.WithField("stack", string(debug.Stack())).Warn("Incorrect logging level")
	}
}
