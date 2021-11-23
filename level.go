package logs

import (
	"encoding/json"
	"fmt"
)

const (
	LevelCritical = "CRITICAL"
	LevelError    = "ERROR"
	LevelWarn     = "WARN"
	LevelInfo     = "INFO"
	LevelDebug    = "DEBUG"
	LevelTrace    = "TRACE"
)

// Level represents a log level
type Level string

func (l Level) Validate() (err error) {
	switch l {
	case LevelCritical:
	case LevelError:
	case LevelWarn:
	case LevelInfo:
	case LevelDebug:
	case LevelTrace:

	default:
		return fmt.Errorf("invalid level, <%s> is not supported", l)
	}

	return nil
}

func (l *Level) UnmarshalJSON(bs []byte) (err error) {
	var str string
	if err = json.Unmarshal(bs, &str); err != nil {
		return
	}

	level := Level(str)

	if err = level.Validate(); err != nil {
		return
	}

	*l = level
	return
}

func (l Level) MarshalJSON() (bs []byte, err error) {
	if err = l.Validate(); err != nil {
		return
	}

	return json.Marshal(string(l))
}
