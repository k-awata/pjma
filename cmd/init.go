/*
Copyright © 2022 K.Awata

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const defaultConf = `# pjma 1.0.0 env file
apps:
  e3d: C:\Program Files (x86)\AVEVA\Everything3D3.1\launch.bat
  adm: C:\Program Files (x86)\AVEVA\Administration1.9\admin.bat
  new: C:\Program Files (x86)\AVEVA\Administration1.9\projectcreation.bat
context:
  module: ""
  tty: false
  project: ""
  user: ""
  password: ""
  mdb: ""
  macro: ""
projects_dir: .\projects
refer_pj:
  - C:\Users\Public\Documents\AVEVA\Projects\E3D3.1\AvevaCatalogue
  - C:\Users\Public\Documents\AVEVA\Projects\E3D3.1\AvevaMarineSample
  - C:\Users\Public\Documents\AVEVA\Projects\E3D3.1\AvevaPlantSample
join_env:
  caf_uic_path:
    - .\cafuic
  pmllib:
    - .\pmllib
  pmlui:
    - .\pmlui
after_cmd: |
  cd /d %temp%
scripts:
  setup: cmd /c "mkdir projects cafuic pmllib pmlui"
`

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a pjma.yaml to current directory",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("pjma.yaml"); err == nil {
			cmd.Println("pjma.yaml already exists")
			return
		}
		f, err := os.Create("pjma.yaml")
		cobra.CheckErr(err)
		defer f.Close()
		_, err = f.WriteString(defaultConf)
		cobra.CheckErr(err)
		cmd.Println("pjma created pjma.yaml to current directory")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
