package logger

import (
	"fmt"
	. "github.com/Sirupsen/logrus"
	"github.com/mattn/go-colorable"
	"time"
)

var g_defaultFormatter Formatter

func init() {
	SetDefaultFormatter()
}

func SetColoredFormatter() {
	SetFormatter(&TextFormatter{ForceColors: true})
	SetOutput(colorable.NewColorableStdout())
}

func EnableDebugging() {
	SetLevel(DebugLevel)
}

func SetDefaultFormatter() {
	if nil == g_defaultFormatter {
		g_defaultFormatter = new(DefaultFormatter)
	}
	SetFormatter(g_defaultFormatter)
}

func GetComponentLogger(component string) *Entry {
	return WithField("component", component)
}

type DefaultFormatter struct{}

func (f *DefaultFormatter) Format(entry *Entry) ([]byte, error) {
	comp, _ := entry.Data["component"].(string)
	return []byte(fmt.Sprintf("[%s] %s _%s_ %v\n", entry.Level, entry.Time.Format(time.RFC3339), comp, entry.Message)), nil

}
