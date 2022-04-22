package ex311

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	res := bytes.Buffer{}
	mStart := 0
	if s[0] == '-' || s[0] == '+' {
		res.WriteByte(s[0])
		mStart = 1
	}

	mEnd := strings.Index(s, ".")
	if mEnd == -1 {
		mEnd = len(s)
	}

	m := s[mStart:mEnd]
	pre := len(m) % 3

	if pre > 0 {
		res.Write([]byte(m[:pre]))
		if len(m) > pre {
			res.WriteByte(',')
		}
	}
	for i, c := range m[pre:] {
		if i%3 == 0 && i != 0 {
			res.WriteByte(',')
		}
		res.WriteRune(c)
	}
	res.WriteString(s[mEnd:])
	return res.String()

}
