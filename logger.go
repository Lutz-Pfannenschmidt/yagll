package yagll

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Logger struct {
	colors     map[Level]string
	files      map[Level]*os.File
	timeFormat string
}

func (l *Logger) SetColor(level Level, color string) {
	l.colors[level] = color
}

func (l *Logger) SetDefaultColors() {
	l.colors[DEBUG] = Blue
	l.colors[INFO] = Green
	l.colors[ERROR] = Red
}

func (l *Logger) SetOutput(level Level, file *os.File) {
	l.files[level] = file
}

func (l *Logger) SetDefaultOutput() {
	l.files[DEBUG] = os.Stdout
	l.files[INFO] = os.Stdout
	l.files[ERROR] = os.Stderr
}

func (l *Logger) SetTimeFormat(format string) {
	l.timeFormat = format
}

func (l *Logger) SetDefaultTimeFormat() {
	l.timeFormat = "Jan _2 15:04:05 2006"
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	l.output(fmt.Sprintf(format, a...), DEBUG)
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.output(fmt.Sprintf(format, a...), INFO)
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.output(fmt.Sprintf(format, a...), ERROR)
}

func (l *Logger) Debugln(message string) {
	msg := strings.TrimRight(message, "\n") + "\n"
	l.output(msg, DEBUG)
}

func (l *Logger) Infoln(message string) {
	msg := strings.TrimRight(message, "\n") + "\n"
	l.output(msg, INFO)
}

func (l *Logger) Errorln(message string) {
	msg := strings.TrimRight(message, "\n") + "\n"
	l.output(msg, ERROR)
}

func NewLogger() *Logger {
	l := &Logger{}
	l.colors = make(map[Level]string)
	l.SetDefaultColors()
	return l
}

func (l *Logger) output(message string, level Level) {
	t := time.Now()
	fmt.Print(t.Format(l.timeFormat))
	fmt.Print(" ")
	fmt.Printf("%s%s%s\n", l.colors[level], message, Reset)
}
