package logs

import (
	"encoding/json"
	"fmt"
)

type Log struct {
	logCore

	Message string `json:"message"`
	Level   Level  `json:"level,omitempty"`
}

func (l *Log) String() string {
	var (
		bs  []byte
		err error
	)

	if bs, err = json.Marshal(l); err != nil {
		fmt.Printf("error marshaling datadog log: %v\n", err)
		return ""
	}

	return string(bs)
}
