package yagll

import (
	"fmt"
	"strings"
)

type Level int

const (
	INFO Level = iota
	DEBUG
	ERROR
)

var levelStrings = map[Level]string{
	INFO:  "INFO",
	DEBUG: "DEBUG",
	ERROR: "ERROR",
}

// Returns the string representation of the level.
func (l Level) String() string {
	max_width := 0
	for _, v := range levelStrings {
		if len(v) > max_width {
			max_width = len(v)
		}
	}
	return fmt.Sprintf("%s%s", levelStrings[l], strings.Repeat(" ", max_width-len(levelStrings[l])))
}
