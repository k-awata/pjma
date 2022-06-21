package pjma

import (
	"reflect"
	"testing"
)

func TestParseCommand(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"normal", args{`cmd /c arg root\path\file`}, []string{`cmd`, `/c`, `arg`, `root\path\file`}},
		{"quote", args{`cmd /c "arg root\path\file"`}, []string{`cmd`, `/c`, `arg root\path\file`}},
		{"single in double", args{`cmd /c "arg 'root\test path\file'"`}, []string{`cmd`, `/c`, `arg 'root\test path\file'`}},
		{"double in single", args{`cmd /c 'arg "root\path\file"'`}, []string{`cmd`, `/c`, `arg "root\path\file"`}},
		{"double in double", args{`cmd /c "arg \"root\path\file\""`}, []string{`cmd`, `/c`, `arg "root\path\file"`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCommand(tt.args.cmd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
