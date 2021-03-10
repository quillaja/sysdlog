// Package sysdlog provides a simple wrapper around a log.Logger to add
// log level prefixes for the systemd log levels.
package sysdlog

import (
	"fmt"
	"log"
)

// Level is a log level.
type Level uint8

const (
	// system is unusable
	Emerg Level = iota
	// action must be taken immediately
	Alert
	// critical conditions
	Crit
	// error conditions
	Err
	// warning conditions
	Warning
	// normal but significant condition
	Notice
	// informational
	Info
	// debug-level messages
	Debug
)

// string versions of levels
var lvlNames = []string{
	"EMERG",   // system is unusable
	"ALERT",   // action must be taken immediately
	"CRIT",    // critical conditions
	"ERR",     // error conditions
	"WARNING", // warning conditions
	"NOTICE",  // normal but significant condition
	"INFO",    // informational
	"DEBUG",   // debug-level messages
}

// LevelLogger wraps a log.Logger.
type LevelLogger struct {
	*log.Logger
	showName bool
}

func NewLevelLogger(logger *log.Logger) *LevelLogger {
	return &LevelLogger{logger, false}
}

// ShowName toggles the logger to add a human readable
// level name to the log message prefix.
func (l *LevelLogger) ShowName(show bool) {
	l.showName = show
}

// SetLevel sets the log level. The level persists until
// it is changed by another call to SetLevel.
func (l *LevelLogger) SetLevel(lvl Level) {
	const (
		noname  = "<%d>"
		yesname = "<%d>%s "
	)
	var newprefix string
	if l.showName {
		newprefix = fmt.Sprintf(yesname, lvl, lvlNames[lvl])
	} else {
		newprefix = fmt.Sprintf(noname, lvl)
	}
	l.Logger.SetPrefix(newprefix)
}
