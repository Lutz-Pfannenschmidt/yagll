package yagll

import "os"

var stdlogger *Logger = NewLogger()

// Sets the color used by the logger for the given level.
func SetColor(level Level, color string) {
	stdlogger.SetColor(level, color)
}

// Set output file for the given level.
func SetOutput(level Level, file *os.File) {
	stdlogger.SetOutput(level, file)
}

// Sets the time format used by the logger (time.Now().Format(format)).
func SetTimeFormat(format string) {
	stdlogger.SetTimeFormat(format)
}

// Sets the default colors used by the logger.
func SetDefaultColors() {
	stdlogger.SetDefaultColors()
}

// Sets the default output files used by the logger (Stdout for DEBUG and INFO, Stderr for ERROR).
func SetDefaultOutput() {
	stdlogger.SetDefaultOutput()
}

// Sets the time format to the default format used by the logger.
func SetDefaultTimeFormat() {
	stdlogger.SetDefaultTimeFormat()
}

// Print the formatted message to the debug level output.
// The message is formatted using fmt.Sprintf and always ends in a new line.
func Debugf(format string, a ...interface{}) {
	stdlogger.Debugf(format, a...)
}

// Print the formatted message to the info level output.
// The message is formatted using fmt.Sprintf and always ends in a new line.
func Infof(format string, a ...interface{}) {
	stdlogger.Infof(format, a...)
}

// Print the formatted message to the error level output.
// The message is formatted using fmt.Sprintf and always ends in a new line.
func Errorf(format string, a ...interface{}) {
	stdlogger.Errorf(format, a...)
}

// Print the message to the debug level output.
// The message always ends in a new line.
func Debugln(message string) {
	stdlogger.Debugln(message)
}

// Print the message to the info level output.
// The message always ends in a new line.
func Infoln(message string) {
	stdlogger.Infoln(message)
}

// Print the message to the error level output.
// The message always ends in a new line.
func Errorln(message string) {
	stdlogger.Errorln(message)
}

// Toggle the output for the given level.
func Toggle(level Level, toggle bool) {
	stdlogger.Toggle(level, toggle)
}

// Sets the default template used by the logger.
// text/template is used to format the output.
func SetDefaultTemplate() {
	stdlogger.SetDefaultTemplate()
}

// Sets the template used by the logger.
// text/template is used to format the output.
// Default template:
// {{ .Time }} [{{ .Color }}{{ .Level }}{{ .Reset }}] {{ .Message }}
//
// The following variables are available:
//
// .Time: The current time formatted using the time format.
// Format can be set using SetTimeFormat.
//
// .Color: The ANSI color for the level.
//
// .Level: The level as string.
//
// .Reset: The ANSI reset code.
//
// .Message: The message.
func SetTemplate(template string) {
	stdlogger.SetTemplate(template)
}
