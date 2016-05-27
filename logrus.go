package logger

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"time"
)

var g_defaultFormatter logrus.Formatter

func init() {
	SetDefaultFormatter()
}

func EnableDebugging() {
	logrus.SetLevel(logrus.DebugLevel)
}

func SetDefaultFormatter() {
	if nil == g_defaultFormatter {
		g_defaultFormatter = new(DefaultFormatter)
	}
	logrus.SetFormatter(g_defaultFormatter)
}

func GetComponentLogger(component string) *logrus.Entry {
	return logrus.WithField("component", component)
}

type DefaultFormatter struct{}

func (f *DefaultFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	comp, _ := entry.Data["component"].(string)
	return []byte(fmt.Sprintf("[%s] %s _%s_ %v\n", entry.Level, entry.Time.Format(time.RFC3339Nano), comp, entry.Message)), nil

}
