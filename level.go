package yagll

type Level int

const (
	INFO Level = iota
	DEBUG
	ERROR
)

// Level to string
func (l Level) String(color string) string {
	switch l {
	case INFO:
		return "[" + color + "INFO " + Reset + "]"
	case DEBUG:
		return "[" + color + "DEBUG" + Reset + "]"
	case ERROR:
		return "[" + color + "ERROR" + Reset + "]"
	}
	return "[" + color + "  ?  " + Reset + "]"
}
