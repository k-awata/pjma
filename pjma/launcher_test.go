package pjma

import (
	"testing"
)

func TestLauncher_MakeBat(t *testing.T) {
	tests := []struct {
		name string
		l    *Launcher
		want string
	}{
		{
			"abs",
			NewLauncher(`c:\root\projects`, `launch1.bat`, "app1"),
			`@echo off
set projects_dir=c:\root\projects
start "app1" /wait cmd /c "launch1.bat"
`,
		},
		{
			"rel",
			NewLauncher(`projects`, `launch2.bat`, "app2"),
			`@echo off
set projects_dir=%cd%\projects
start "app2" /wait cmd /c "launch2.bat"
`,
		},
		{
			"module",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetModule("Design"),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat" Design
`,
		},
		{
			"tty",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetTty(true),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat" TTY
`,
		},
		{
			"pj",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetProject("AAA"),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat"
`,
		},
		{
			"up",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetUser("SYSTEM", "XXXXXX"),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat"
`,
		},
		{
			"pju",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetProject("AAA").
				SetUser("SYSTEM", ""),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat"
`,
		},
		{
			"pjp",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetProject("AAA").
				SetUser("", "XXXXXX"),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat"
`,
		},
		{
			"pjup",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetProject("AAA").
				SetUser("SYSTEM", "XXXXXX"),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat" AAA SYSTEM/XXXXXX
`,
		},
		{
			"mdb only",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetMdb("ALL"),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat"
`,
		},
		{
			"mdb",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetProject("AAA").
				SetUser("SYSTEM", "XXXXXX").
				SetMdb("ALL"),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat" AAA SYSTEM/XXXXXX /ALL
`,
		},
		{
			"slash mdb",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetProject("AAA").
				SetUser("SYSTEM", "XXXXXX").
				SetMdb("/ALL"),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat" AAA SYSTEM/XXXXXX /ALL
`,
		},
		{
			"macro only",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetMacro("Draw"),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat"
`,
		},
		{
			"macro",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetProject("AAA").
				SetUser("SYSTEM", "XXXXXX").
				SetMacro("Design"),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat" AAA SYSTEM/XXXXXX Design
`,
		},
		{
			"all",
			NewLauncher(`projects`, `launch.bat`, "app").
				SetModule("Monitor").
				SetTty(true).
				SetProject("AAA").
				SetUser("SYSTEM", "XXXXXX").
				SetMdb("ALL").
				SetMacro("Draw"),
			`@echo off
set projects_dir=%cd%\projects
start "app" /wait cmd /c "launch.bat" Monitor TTY AAA SYSTEM/XXXXXX /ALL Draw
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.MakeBat(); got != tt.want {
				t.Errorf("Launcher.MakeBat() = %v, want %v", got, tt.want)
			}
		})
	}
}
