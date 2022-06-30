package pjma

import "testing"

func TestEvars_makeJoinEnv(t *testing.T) {
	tests := []struct {
		name string
		e    *Evars
		want string
	}{
		{
			"test",
			&Evars{
				jenv: map[string][]string{
					"caf_uic_path": {
						`c:\root\cafuic_1`,
						`c:\root\cafuic_2`,
						`cafuic_3`,
					},
					"pmllib": {
						`c:\root\pmllib_1`,
						`c:\root\pmllib_2`,
						`pmllib_3`,
					},
					"pmlui": {
						`c:\root\pmlui_1`,
						`c:\root\pmlui_2`,
						`pmlui_3`,
					},
				},
			},
			`set caf_uic_path=c:\root\cafuic_1;c:\root\cafuic_2;%cd%\cafuic_3;%caf_uic_path%
set pmllib=c:\root\pmllib_1;c:\root\pmllib_2;%cd%\pmllib_3;%pmllib%
set pmlui=c:\root\pmlui_1;c:\root\pmlui_2;%cd%\pmlui_3;%pmlui%

`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.makeJoinEnv(); got != tt.want {
				t.Errorf("Evars.makeJoinEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
