package pjma

import (
	"bytes"
	"sort"

	"golang.org/x/text/encoding/htmlindex"
	"golang.org/x/text/transform"
)

// SortStringKeys returns sorted keys from string map
func SortStringKeys[T any](m map[string]T) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// MergeStringMaps appends keys and values of another map to a map but skips for existing key
func MergeStringMaps(src map[string]interface{}, add map[string]interface{}) map[string]interface{} {
	for k, v := range add {
		if _, ok := src[k]; !ok {
			src[k] = v
		}
	}
	return src
}

// ParseCommand returns command and arguments from string
func ParseCommand(cmd string) []string {
	var buf bytes.Buffer
	var quote rune
	param := []string{}
	esc := false
	for _, r := range cmd {
		// start to escape
		if !esc && r == '\\' {
			esc = true
			continue
		}
		// Escaping
		if esc {
			if r != '"' && r != '\'' && r != '\\' {
				buf.WriteRune('\\')
			}
			buf.WriteRune(r)
			esc = false
			continue
		}
		// Find end quote
		if r == quote {
			quote = 0
			continue
		}
		// Find begin quote
		if quote == 0 && (r == '"' || r == '\'') {
			quote = r
			continue
		}
		// Cut command by space
		if quote == 0 && r == ' ' {
			if buf.Len() > 0 {
				param = append(param, buf.String())
				buf.Reset()
			}
			continue
		}
		buf.WriteRune(r)
	}
	return append(param, buf.String())
}

// EncodeForBatch returns appropriate format string for batch files in specified encoding
func EncodeForBatch(s string, encode string) (string, error) {
	e, err := htmlindex.Get(encode)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	w := transform.NewWriter(&buf, e.NewEncoder())
	if _, err := w.Write(bytes.ReplaceAll([]byte(s), []byte("\n"), []byte("\r\n"))); err != nil {
		return "", err
	}
	return buf.String(), nil
}
