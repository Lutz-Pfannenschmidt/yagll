package yagll

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

type Logger struct {
	colors     map[Level]string
	files      map[Level]*os.File
	toggle     map[Level]bool
	template   string
	timeFormat string
}

func (l *Logger) SetColor(level Level, color string) {
	l.colors[level] = color
}

func (l *Logger) SetDefaultColors() {
	l.colors[DEBUG] = Yellow
	l.colors[INFO] = Blue
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

// Sets the template used by the logger.
// Uses text/template to format the message.
func (l *Logger) SetTemplate(template string) {
	l.template = template
}

// Sets the template to the default template used by the logger.
// Uses text/template to format the message.
//
// Default template: {{ .Time }} [{{ .Color }}{{ .Level }}{{ .Reset }}] {{.Message}}
func (l *Logger) SetDefaultTemplate() {
	l.template = `{{ .Time }} [{{ .Color }}{{ .Level }}{{ .Reset }}] {{.Message}}`
}

// Sets the time format used by the logger (time.Now().Format(format)).
func (l *Logger) SetTimeFormat(format string) {
	l.timeFormat = format
}

// Sets the time format to the default format used by the logger.
func (l *Logger) SetDefaultTimeFormat() {
	l.timeFormat = "Jan _2 15:04:05 2006"
}

// Print the formatted message to the debug level output.
// The message is formatted using fmt.Sprintf and always ends in a new line.
func (l *Logger) Debugf(format string, a ...interface{}) {
	l.output(fmt.Sprintf(format, a...), DEBUG)
}

// Print the formatted message to the info level output.
// The message is formatted using fmt.Sprintf and always ends in a new line.
func (l *Logger) Infof(format string, a ...interface{}) {
	l.output(fmt.Sprintf(format, a...), INFO)
}

// Print the formatted message to the error level output.
// The message is formatted using fmt.Sprintf and always ends in a new line.
func (l *Logger) Errorf(format string, a ...interface{}) {
	l.output(fmt.Sprintf(format, a...), ERROR)
}

// Print the message to the debug level output.
// The message always ends in a new line.
func (l *Logger) Debugln(message string) {
	l.output(message, DEBUG)
}

// Print the message to the info level output.
// The message always ends in a new line.
func (l *Logger) Infoln(message string) {
	l.output(message, INFO)
}

// Print the message to the error level output.
// The message always ends in a new line.
func (l *Logger) Errorln(message string) {
	l.output(message, ERROR)
}

// Toggle weather or not to print a message to the output.
func (l *Logger) Toggle(level Level, toggle bool) {
	l.toggle[level] = toggle
}

func NewLogger() *Logger {
	l := &Logger{}
	l.colors = make(map[Level]string)
	l.files = make(map[Level]*os.File)
	l.toggle = map[Level]bool{DEBUG: true, INFO: true, ERROR: true}
	l.SetDefaultTemplate()
	l.SetDefaultColors()
	l.SetDefaultOutput()
	l.SetDefaultTimeFormat()
	return l
}

func (l *Logger) output(message string, level Level) {
	if !l.toggle[level] {
		return
	}
	t := time.Now()
	tstring := t.Format(l.timeFormat)
	msg := strings.TrimRight(message, "\n")
	template, err := template.New("yagll").Parse(l.template)
	if err != nil {
		fmt.Fprintf(os.Stderr, "YAGLL: Error parsing template: %s", err)
	}
	err = template.Execute(l.files[level], map[string]string{
		"Time":    tstring,
		"Color":   l.colors[level],
		"Level":   level.String(),
		"Reset":   Reset,
		"Message": msg,
	})
	l.files[level].Write([]byte("\n"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "YAGLL: Error executing template: %s", err)
	}
}
