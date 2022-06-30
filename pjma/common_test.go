package pjma

import (
	"reflect"
	"testing"
)

func TestSortStringKeys(t *testing.T) {
	type args struct {
		m map[string]any
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test",
			args{map[string]any{
				"he":    struct{}{},
				"dalet": struct{}{},
				"gimel": struct{}{},
				"bet":   struct{}{},
				"alef":  struct{}{},
			}},
			[]string{"alef", "bet", "dalet", "gimel", "he"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortStringKeys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortStringKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeStringMaps(t *testing.T) {
	type args struct {
		src map[string]interface{}
		add map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			"test",
			args{
				map[string]interface{}{
					"alef":  "src",
					"bet":   "src",
					"gimel": "src",
					"dalet": "src",
					"he":    "src",
					"vav":   "src",
				},
				map[string]interface{}{
					"dalet": "add",
					"he":    "add",
					"vav":   "add",
					"zayin": "add",
					"chet":  "add",
					"tet":   "add",
				},
			},
			map[string]interface{}{
				"alef":  "src",
				"bet":   "src",
				"gimel": "src",
				"dalet": "src",
				"he":    "src",
				"vav":   "src",
				"zayin": "add",
				"chet":  "add",
				"tet":   "add",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeStringMaps(tt.args.src, tt.args.add); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeStringMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCommand(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"normal", args{`cmd /c arg root\path\file`}, []string{`cmd`, `/c`, `arg`, `root\path\file`}},
		{"quote", args{`cmd /c "arg root\path\file"`}, []string{`cmd`, `/c`, `arg root\path\file`}},
		{"single in double", args{`cmd /c "arg 'root\test path\file'"`}, []string{`cmd`, `/c`, `arg 'root\test path\file'`}},
		{"double in single", args{`cmd /c 'arg "root\test path\file"'`}, []string{`cmd`, `/c`, `arg "root\test path\file"`}},
		{"double in double", args{`cmd /c "arg \"root\test path\file\""`}, []string{`cmd`, `/c`, `arg "root\test path\file"`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCommand(tt.args.cmd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
