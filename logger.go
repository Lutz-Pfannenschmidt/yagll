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

// Sets the time format used by the logger (time.Now().Format(format))
func (l *Logger) SetTimeFormat(format string) {
	l.timeFormat = format
}

// Sets the time format to the default format used by the logger
func (l *Logger) SetDefaultTimeFormat() {
	l.timeFormat = "Jan _2 15:04:05 2006"
}

// Print the formatted message to the debug level output
// The message is formatted using fmt.Sprintf and always ends in a new line
func (l *Logger) Debugf(format string, a ...interface{}) {
	l.output(fmt.Sprintf(format, a...), DEBUG)
}

// Print the formatted message to the info level output
// The message is formatted using fmt.Sprintf and always ends in a new line
func (l *Logger) Infof(format string, a ...interface{}) {
	l.output(fmt.Sprintf(format, a...), INFO)
}

// Print the formatted message to the error level output
// The message is formatted using fmt.Sprintf and always ends in a new line
func (l *Logger) Errorf(format string, a ...interface{}) {
	l.output(fmt.Sprintf(format, a...), ERROR)
}

// Print the message to the debug level output
// The message always ends in a new line
func (l *Logger) Debugln(message string) {
	l.output(message, DEBUG)
}

// Print the message to the info level output
// The message always ends in a new line
func (l *Logger) Infoln(message string) {
	l.output(message, INFO)
}

// Print the message to the error level output
// The message always ends in a new line
func (l *Logger) Errorln(message string) {
	l.output(message, ERROR)
}

func NewLogger() *Logger {
	l := &Logger{}
	l.colors = make(map[Level]string)
	l.files = make(map[Level]*os.File)
	l.SetDefaultColors()
	l.SetDefaultOutput()
	l.SetDefaultTimeFormat()
	return l
}

func (l *Logger) output(message string, level Level) {
	t := time.Now()
	fmt.Print(t.Format(l.timeFormat))
	fmt.Print(" ")
	msg := strings.TrimRight(message, "\n")
	fmt.Printf("%s%s%s\n", l.colors[level], msg, Reset)
}
