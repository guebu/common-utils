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
	LogLevelEnv = "LOG_LEVEL_GO"
)


// ToDo: Why is somewhere log2.blabla... being used!!
// Try out something new...
var(
	Log *myLogger
)

type myLogger struct {
	log *logrus.Logger
}

func init(){

	var tmpLogger = &logrus.Logger{
		Level: getLogLevel(),
		Out:os.Stdout,
	}

	Log = &myLogger{}
	Log.log = tmpLogger

	if IsProduction() {
		// In Production we use the JSON output
		Log.log.Formatter = &logrus.JSONFormatter{}
	} else {
		//Log.Formatter = &logrus.JSONFormatter{}
		Log.log.Formatter = &logrus.TextFormatter{}
	}
}

func getLogLevel( ) logrus.Level {
	switch os.Getenv(LogLevelEnv){
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}

func (log *myLogger)Printf(format string, v...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}

}

func Info(msg string, tags ...string) {
	if Log.log.Level < logrus.InfoLevel {
		return
	}
	Log.log.WithFields(parseFields(tags...)).Info(msg)
}

func Error(msg string, err error, tags ...string) {
	if Log.log.Level < logrus.ErrorLevel {
		return
	}
	msg = fmt.Sprintf("%s - ERROR: -- %v", msg, err)
	Log.log.WithFields(parseFields(tags...)).Error(msg)
}

func Debug(msg string, tags ...string) {
	if Log.log.Level < logrus.DebugLevel {
		return
	}
	Log.log.WithFields(parseFields(tags...)).Debug(msg)
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