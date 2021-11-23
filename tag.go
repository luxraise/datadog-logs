package logs

import (
	"encoding/json"
	"fmt"
	"strings"
)

func makeTagFromString(str string) (t Tag, err error) {
	spl := strings.Split(str, ":")
	if len(spl) < 2 {
		err = fmt.Errorf("invalid parts, expected 2 and received %d", len(spl))
		return
	}

	t.Key = spl[0]
	t.Value = spl[1]
	return
}

// Tag represents a log tag
// e.g. env:staging, version:5.1, etc
type Tag struct {
	Key   string
	Value string
}

func (t *Tag) UnmarshalJSON(bs []byte) (err error) {
	var str string
	if err = json.Unmarshal(bs, &str); err != nil {
		return
	}

	var tag Tag
	if tag, err = makeTagFromString(str); err != nil {
		return
	}

	*t = tag
	return
}

func (t Tag) MarshalJSON() (bs []byte, err error) {
	return json.Marshal(t.String())
}

func (t Tag) String() string {
	return t.Key + ":" + t.Value
}
