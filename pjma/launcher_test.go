package pjma

import (
	"reflect"
	"testing"
)

func TestLauncher_SetMdb(t *testing.T) {
	type args struct {
		m string
	}
	tests := []struct {
		name string
		l    *Launcher
		args args
		want *Launcher
	}{
		// TODO: Add test cases.
		{"Empty", &Launcher{}, args{""}, &Launcher{}},
		{"No slash", &Launcher{}, args{"MDB"}, &Launcher{mdb: "/MDB"}},
		{"slash", &Launcher{}, args{"/MDB"}, &Launcher{mdb: "/MDB"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.SetMdb(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Launcher.SetMdb() = %v, want %v", got, tt.want)
			}
		})
	}
}
