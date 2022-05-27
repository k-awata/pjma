package pjma

import (
	"bytes"
)

func ParseScript(scr string) []string {
	var buf bytes.Buffer
	var quote rune
	params := []string{}
	esc := false
	for _, r := range scr {
		if !esc && r == '\\' {
			esc = true
			continue
		}
		if esc {
			if r != '"' && r != '\'' && r != '\\' {
				buf.WriteRune('\\')
			}
			buf.WriteRune(r)
			esc = false
			continue
		}
		if r == quote {
			quote = 0
			continue
		}
		if quote == 0 && (r == '"' || r == '\'') {
			quote = r
			continue
		}
		if quote == 0 && r == ' ' {
			if buf.Len() > 0 {
				params = append(params, buf.String())
				buf.Reset()
			}
			continue
		}
		buf.WriteRune(r)
	}
	return append(params, buf.String())
}
