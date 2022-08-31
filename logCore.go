package logs

import "fmt"

type logCore struct {
	Hostname string `json:"hostname,omitempty"`
	Service  string `json:"service,omitempty"`

	Tags Tags `json:"ddtags,omitempty"`
}

func (c logCore) Log(message string, tags ...Tag) {
	c.log(message, LevelInfo, nil, tags)
}

func (c logCore) Warn(message string, tags ...Tag) {
	c.log(message, LevelWarn, nil, tags)
}

func (c logCore) Error(message string, tags ...Tag) {
	c.log(message, LevelError, nil, tags)
}

func (c logCore) Critical(message string, tags ...Tag) {
	c.log(message, LevelCritical, nil, tags)
}

func (c logCore) Debug(message string, tags ...Tag) {
	c.log(message, LevelDebug, nil, tags)
}

func (c logCore) Trace(message string, tags ...Tag) {
	c.log(message, LevelTrace, nil, tags)
}

func (c logCore) LogWithData(message string, data interface{}, tags ...Tag) {
	c.log(message, LevelInfo, data, tags)
}

func (c logCore) WarnWithData(message string, data interface{}, tags ...Tag) {
	c.log(message, LevelWarn, data, tags)
}

func (c logCore) ErrorWithData(message string, data interface{}, tags ...Tag) {
	c.log(message, LevelError, data, tags)
}

func (c logCore) CriticalWithData(message string, data interface{}, tags ...Tag) {
	c.log(message, LevelCritical, data, tags)
}

func (c logCore) DebugWithData(message string, data interface{}, tags ...Tag) {
	c.log(message, LevelDebug, data, tags)
}

func (c logCore) TraceWithData(message string, data interface{}, tags ...Tag) {
	c.log(message, LevelTrace, data, tags)
}

func (c logCore) Logf(message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.log(message, LevelInfo, nil, nil)
}

func (c logCore) Warnf(message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.log(message, LevelWarn, nil, nil)
}

func (c logCore) Errorf(message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.log(message, LevelError, nil, nil)
}

func (c logCore) Criticalf(message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.log(message, LevelCritical, nil, nil)
}

func (c logCore) Debugf(message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.log(message, LevelDebug, nil, nil)
}

func (c logCore) Tracef(message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.log(message, LevelTrace, nil, nil)
}

func (c logCore) LogWithDataf(message string, data interface{}, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.log(message, LevelInfo, data, nil)
}

func (c logCore) WarnWithDataf(message string, data interface{}, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.log(message, LevelWarn, data, nil)
}

func (c logCore) ErrorWithDataf(message string, data interface{}, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.log(message, LevelError, data, nil)
}

func (c logCore) CriticalWithDataf(message string, data interface{}, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.log(message, LevelCritical, data, nil)
}

func (c logCore) DebugWithDataf(message string, data interface{}, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.log(message, LevelDebug, data, nil)
}

func (c logCore) TraceWithDataf(message string, data interface{}, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.log(message, LevelTrace, data, nil)
}

func (c logCore) log(message string, level Level, data interface{}, tags []Tag) {
	log := c.makeLog(message, level, data, tags)
	msg := log.String()
	fmt.Println(msg)
}

func (c logCore) copy() (copied logCore) {
	copied = c
	copied.Tags = make(Tags, len(c.Tags))
	copy(copied.Tags, c.Tags)
	return
}

func (c logCore) makeLog(message string, level Level, data interface{}, tags Tags) (log Log) {
	log.logCore = c.copy()
	log.Level = LevelInfo
	log.Message = message
	log.Data = data
	log.Tags = append(log.Tags, tags...)
	return
}
