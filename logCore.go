package logs

import "fmt"

type logCore struct {
	Hostname string `json:"hostname"`
	Source   string `json:"ddsource"`
	Service  string `json:"service"`
}

func (c logCore) Log(message string, tags ...Tag) {
	log := c.makeLog(message, tags, LevelInfo)
	msg := log.String()
	fmt.Println(msg)
}

func (c logCore) Warn(message string, tags ...Tag) {
	log := c.makeLog(message, tags, LevelInfo)
	msg := log.String()
	fmt.Println(msg)
}

func (c logCore) Error(message string, tags ...Tag) {
	log := c.makeLog(message, tags, LevelError)
	msg := log.String()
	fmt.Println(msg)
}

func (c logCore) Critical(message string, tags ...Tag) {
	log := c.makeLog(message, tags, LevelCritical)
	msg := log.String()
	fmt.Println(msg)
}

func (c logCore) Debug(message string, tags ...Tag) {
	log := c.makeLog(message, tags, LevelDebug)
	msg := log.String()
	fmt.Println(msg)
}

func (c logCore) Trace(message string, tags ...Tag) {
	log := c.makeLog(message, tags, LevelTrace)
	msg := log.String()
	fmt.Println(msg)
}

func (c logCore) makeLog(message string, tags Tags, level Level) (log Log) {
	log.logCore = c
	log.Tags = tags
	log.Level = LevelInfo
	return
}
