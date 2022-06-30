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
	"fmt"
	"strings"

	"github.com/k-awata/pjma/pjma"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// mkbatCmd represents the mkbat command
var mkbatCmd = &cobra.Command{
	Use:     "mkbat app_name",
	Short:   "Output bat file commands to launch an app",
	Example: `  pjma mkbat e3d > launch.bat`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bindContextFlags(cmd)
		batkey := "apps." + args[0]
		if !viper.IsSet(batkey) {
			cobra.CheckErr("pjma cannot find bat name: " + args[0])
		}

		lnchr := pjma.NewLauncher(viper.GetString("projects_dir"), viper.GetString(batkey), args[0]).
			SetModule(viper.GetString("context.module")).
			SetTty(viper.GetBool("context.tty")).
			SetProject(viper.GetString("context.project")).
			SetUser(viper.GetString("context.user"), viper.GetString("context.password")).
			SetMdb(viper.GetString("context.mdb")).
			SetMacro(viper.GetString("context.macro"))

		fmt.Print(strings.ReplaceAll(lnchr.MakeBat(), "\n", "\r\n"))
	},
}

func init() {
	rootCmd.AddCommand(mkbatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mkbatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mkbatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	mkbatCmd.Flags().StringP("module", "e", "", "module to enter")
	mkbatCmd.Flags().BoolP("tty", "t", false, "launch with no gui mode")
	mkbatCmd.Flags().StringP("project", "P", "", "project code")
	mkbatCmd.Flags().StringP("user", "u", "", "username")
	mkbatCmd.Flags().StringP("password", "p", "", "password")
	mkbatCmd.Flags().StringP("mdb", "M", "", "MDB name")
	mkbatCmd.Flags().StringP("macro", "m", "", "macro file")
}
