package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	TagSeparator = ":"
	LogLevel = "info"
	Production = "prod"
)


// ToDo: Why is somewhere log2.blabla... being used!!
var(
	Log *logrus.Logger
)

func init(){
	level, err := logrus.ParseLevel(LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}

	Log = &logrus.Logger{
		Level: level,
		Out:os.Stdout,
	}

	if IsProduction() {
		// In Production we use the JSON output
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		//Log.Formatter = &logrus.JSONFormatter{}
		Log.Formatter = &logrus.TextFormatter{}
	}
}

func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Info(msg)
}

func Error(msg string, err error, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	msg = fmt.Sprintf("%s - ERROR: -- %v", msg, err)
	Log.WithFields(parseFields(tags...)).Error(msg)
}

func Debug(msg string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Debug(msg)
}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		elements := strings.Split(tag, TagSeparator)
		result[strings.TrimSpace(elements[0])] = strings.TrimSpace(elements[1])
	}
	return result
}

func IsProduction() bool {
	return os.Getenv("GO_ENVIRONMENT") == Production
}