/*
Copyright © 2021 K.Awata

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
	"fmt"
	"os"

	"github.com/k-awata/pjma/internal/pjma"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:     "run APPNAME [PJDIR]",
	Aliases: []string{"r"},
	Short:   "Run a project with an app defined by config file",
	Long:    "Run a project with an app defined by config file",
	Example: `pjma run e3d31 -o Design -p APS -u SYSTEM/XXXXXX -d ALL`,
	Args:    cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		// Set app name from 1st arg
		if !viper.IsSet("apps." + args[0]) {
			fmt.Fprintln(os.Stderr, "app name not found in config file")
			return
		}
		viper.Set("context.bat", args[0])

		// Set context from flags
		viper.BindPFlag("context.module", cmd.Flags().Lookup("module"))
		viper.BindPFlag("context.tty", cmd.Flags().Lookup("tty"))
		viper.BindPFlag("context.project", cmd.Flags().Lookup("project"))
		viper.BindPFlag("context.user", cmd.Flags().Lookup("user"))
		viper.BindPFlag("context.mdb", cmd.Flags().Lookup("mdb"))
		viper.BindPFlag("context.macro", cmd.Flags().Lookup("macro"))

		// Directly open project by path from 2nd arg
		if len(args) == 2 {
			pjcode, err := pjma.GetProjectCode(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			viper.Set("extrapj", append(viper.GetStringSlice("extrapj"), args[1]))
			viper.Set("context.project", pjcode)
		}

		if err := pjma.MakeEvars(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if err := pjma.Launch(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringP("module", "o", "", "module to enter")
	runCmd.Flags().BoolP("tty", "t", false, "run with command line only mode")
	runCmd.Flags().StringP("project", "p", "", "project code to open")
	runCmd.Flags().StringP("user", "u", "", "username and password (needs slash between them)")
	runCmd.Flags().StringP("mdb", "d", "", "MDB name")
	runCmd.Flags().StringP("macro", "m", "", "macro file to run")
}
