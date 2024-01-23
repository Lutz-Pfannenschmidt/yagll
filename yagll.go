package yagll

import "os"

var stdlogger *Logger = NewLogger()

func SetColor(level Level, color string) {
	stdlogger.SetColor(level, color)
}

func SetOutput(level Level, file *os.File) {
	stdlogger.SetOutput(level, file)
}

func SetDefaultColors() {
	stdlogger.SetDefaultColors()
}

func SetDefaultOutput() {
	stdlogger.SetDefaultOutput()
}

func Debugf(format string, a ...interface{}) {
	stdlogger.Debugf(format, a...)
}

func Infof(format string, a ...interface{}) {
	stdlogger.Infof(format, a...)
}

func Errorf(format string, a ...interface{}) {
	stdlogger.Errorf(format, a...)
}

func Debugln(message string) {
	stdlogger.Debugln(message)
}

func Infoln(message string) {
	stdlogger.Infoln(message)
}

func Errorln(message string) {
	stdlogger.Errorln(message)
}
