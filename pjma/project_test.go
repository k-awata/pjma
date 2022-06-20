package pjma

import "testing"

func TestProject_DumpEvars(t *testing.T) {
	tests := []struct {
		name string
		p    *Project
		want string
	}{
		// TODO: Add test cases.
		{
			"test",
			&Project{
				`c:\root\projects`,
				`c:\root\projects\TestProject`,
				"tst",
				[]string{
					`TestProject\tst000`,
					`TestProject\tstdflts`,
					`TestProject\tstiso`,
					`TestProject\tstmac`,
					`TestProject\tstpic`,
				},
			},
			`set tst000=%projects_dir%\TestProject\tst000
set tstdflts=%projects_dir%\TestProject\tstdflts
set tstiso=%projects_dir%\TestProject\tstiso
set tstmac=%projects_dir%\TestProject\tstmac
set tstpic=%projects_dir%\TestProject\tstpic
set tst000id=TestProject
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
