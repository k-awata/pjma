package pjma

import (
	"bytes"
	"sort"
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
