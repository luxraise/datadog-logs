package logs

import (
	"encoding/json"
	"strings"
)

func makeTagsFromString(str string) (ts Tags, err error) {
	parts := strings.Split(str, ",")
	ts = make(Tags, 0, len(parts))
	for _, part := range parts {
		var tag Tag
		if tag, err = makeTagFromString(part); err != nil {
			return
		}

		ts = append(ts, tag)
	}

	return
}

type Tags []Tag

func (ts *Tags) UnmarshalJSON(bs []byte) (err error) {
	var str string
	if err = json.Unmarshal(bs, &str); err != nil {
		return
	}

	var tags Tags
	if tags, err = makeTagsFromString(str); err != nil {
		return
	}

	*ts = tags
	return
}

func (t Tags) MarshalJSON() (bs []byte, err error) {
	parts := t.getParts()
	str := strings.Join(parts, ",")
	return json.Marshal(str)
}

func (t Tags) getParts() (parts []string) {
	parts = make([]string, 0, len(t))
	for _, tag := range t {
		parts = append(parts, tag.String())
	}

	return
}
