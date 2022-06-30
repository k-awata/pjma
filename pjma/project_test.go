package pjma

import "testing"

func TestProject_DumpEvars(t *testing.T) {
	tests := []struct {
		name string
		p    *Project
		want string
	}{
		{
			"pjdir",
			&Project{
				`c:\root\projects`,
				`c:\root\projects\AAAProject`,
				"aaa",
				[]string{
					`AAAProject\aaa000`,
					`AAAProject\aaadflts`,
					`AAAProject\aaaiso`,
					`AAAProject\aaamac`,
					`AAAProject\aaapic`,
				},
			},
			`set aaa000=%projects_dir%\AAAProject\aaa000
set aaadflts=%projects_dir%\AAAProject\aaadflts
set aaaiso=%projects_dir%\AAAProject\aaaiso
set aaamac=%projects_dir%\AAAProject\aaamac
set aaapic=%projects_dir%\AAAProject\aaapic
set aaa000id=AAAProject
`,
		},
		{
			"abs",
			&Project{
				``,
				`c:\root\projects\BBBProject`,
				"bbb",
				[]string{
					`c:\root\projects\BBBProject\bbb000`,
					`c:\root\projects\BBBProject\bbbdflts`,
					`c:\root\projects\BBBProject\bbbiso`,
					`c:\root\projects\BBBProject\bbbmac`,
					`c:\root\projects\BBBProject\bbbpic`,
				},
			},
			`set bbb000=c:\root\projects\BBBProject\bbb000
set bbbdflts=c:\root\projects\BBBProject\bbbdflts
set bbbiso=c:\root\projects\BBBProject\bbbiso
set bbbmac=c:\root\projects\BBBProject\bbbmac
set bbbpic=c:\root\projects\BBBProject\bbbpic
set bbb000id=BBBProject
`,
		},
		{
			"rel",
			&Project{
				``,
				`projects\CCCProject`,
				"ccc",
				[]string{
					`projects\CCCProject\ccc000`,
					`projects\CCCProject\cccdflts`,
					`projects\CCCProject\ccciso`,
					`projects\CCCProject\cccmac`,
					`projects\CCCProject\cccpic`,
				},
			},
			`set ccc000=%cd%\projects\CCCProject\ccc000
set cccdflts=%cd%\projects\CCCProject\cccdflts
set ccciso=%cd%\projects\CCCProject\ccciso
set cccmac=%cd%\projects\CCCProject\cccmac
set cccpic=%cd%\projects\CCCProject\cccpic
set ccc000id=CCCProject
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.DumpEvars(); got != tt.want {
				t.Errorf("Project.DumpEvars() = %v, want %v", got, tt.want)
			}
		})
	}
}
