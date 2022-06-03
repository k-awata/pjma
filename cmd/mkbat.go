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

var mkbatCmd = &cobra.Command{
	Use:     "mkbat APPNAME",
	Short:   "Make a batch file to run specified app",
	Long:    "Make a batch file to run specified app",
	Example: `pjma mkbat e3d31 > launch.bat`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Set app name from arg
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
		viper.BindPFlag("absbat", cmd.Flags().Lookup("abs"))

		bat, err := pjma.MakeLaunch(viper.GetBool("absbat"))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Print(bat)
	},
}

func init() {
	rootCmd.AddCommand(mkbatCmd)

	mkbatCmd.Flags().StringP("module", "o", "", "module to enter")
	mkbatCmd.Flags().BoolP("tty", "t", false, "run with command line only mode")
	mkbatCmd.Flags().StringP("project", "p", "", "project code to open")
	mkbatCmd.Flags().StringP("user", "u", "", "username and password (needs slash between them)")
	mkbatCmd.Flags().StringP("mdb", "d", "", "MDB name")
	mkbatCmd.Flags().StringP("macro", "m", "", "macro file to run")
	mkbatCmd.Flags().Bool("abs", false, "set absolute path to projects_dir")
}
