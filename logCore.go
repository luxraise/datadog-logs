package logs

import "fmt"

type logCore struct {
	Hostname string `json:"hostname,omitempty"`
	Source   string `json:"source,omitempty"`
	Service  string `json:"service,omitempty"`

	Tags Tags `json:"ddtags,omitempty"`
}

func (c logCore) Log(message string, tags ...Tag) {
	c.log(message, LevelInfo, tags)
}

func (c logCore) Warn(message string, tags ...Tag) {
	c.log(message, LevelWarn, tags)
}

func (c logCore) Error(message string, tags ...Tag) {
	c.log(message, LevelError, tags)
}

func (c logCore) Critical(message string, tags ...Tag) {
	c.log(message, LevelCritical, tags)
}

func (c logCore) Debug(message string, tags ...Tag) {
	c.log(message, LevelDebug, tags)
}

func (c logCore) Trace(message string, tags ...Tag) {
	c.log(message, LevelTrace, tags)
}

func (c logCore) Logf(message string, args ...interface{}) {
	c.log(message, LevelInfo, nil)
}

func (c logCore) Warnf(message string, args ...interface{}) {
	c.log(message, LevelWarn, nil)
}

func (c logCore) Errorf(message string, args ...interface{}) {
	c.log(message, LevelError, nil)
}

func (c logCore) Criticalf(message string, args ...interface{}) {
	c.log(message, LevelCritical, nil)
}

func (c logCore) Debugf(message string, args ...interface{}) {
	c.log(message, LevelDebug, nil)
}

func (c logCore) Tracef(message string, args ...interface{}) {
	c.log(message, LevelTrace, nil)
}

func (c logCore) log(message string, level Level, tags []Tag) {
	log := c.makeLog(message, tags, level)
	log.Tags = append(log.Tags, tags...)
	msg := log.String()
	fmt.Println(msg)
}

func (c logCore) copy() (copied logCore) {
	copied = c
	copied.Tags = make(Tags, len(c.Tags))
	copy(copied.Tags, c.Tags)
	return
}

func (c logCore) makeLog(message string, tags Tags, level Level) (log Log) {
	log.logCore = c.copy()
	log.Tags = tags
	log.Level = LevelInfo
	log.Message = message
	return
}
