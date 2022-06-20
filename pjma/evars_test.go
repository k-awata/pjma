package pjma

import "testing"

func TestEvars_makeJoinEnv(t *testing.T) {
	tests := []struct {
		name string
		e    *Evars
		want string
	}{
		// TODO: Add test cases.
		{
			"join env",
			&Evars{
				jenv: map[string][]string{
					"caf_uic_path": {
						`c:\cafuic_1`,
						`c:\cafuic_2`,
						`c:\cafuic_3`,
					},
					"pmllib": {
						`c:\pmllib_1`,
						`c:\pmllib_2`,
						`c:\pmllib_3`,
					},
					"pmlui": {
						`c:\pmlui_1`,
						`c:\pmlui_2`,
						`c:\pmlui_3`,
					},
				},
			},
			`set caf_uic_path=c:\cafuic_1;c:\cafuic_2;c:\cafuic_3;%caf_uic_path%
set pmllib=c:\pmllib_1;c:\pmllib_2;c:\pmllib_3;%pmllib%
set pmlui=c:\pmlui_1;c:\pmlui_2;c:\pmlui_3;%pmlui%

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
