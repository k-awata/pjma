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
	"syscall"

	"github.com/k-awata/pjma/pjma"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:     "open app_name",
	Aliases: []string{"o"},
	Short:   "Open an app defined by pjma.yaml",
	Example: `  pjma open e3d -e Design -P APS -u SYSTEM -p XXXXXX -M ALL`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bindContextFlags(cmd)
		batkey := "apps." + args[0]
		if !viper.IsSet(batkey) {
			cobra.CheckErr("pjma cannot find bat name: " + args[0])
		}

		// Save custom_evars.bat
		evars := pjma.NewEvars(viper.GetString("projects_dir"))
		cobra.CheckErr(evars.AddReferProjectDirs(viper.GetStringSlice("refer_pj")))
		evars.AddJoinEnv(viper.GetStringMapStringSlice("join_env"))
		evars.AddAfterCmd(viper.GetString("after_cmd"))

		// Add reference project if given as directory path
		fs, err := os.Stat(viper.GetString("context.project"))
		if err == nil && fs.IsDir() {
			pj, err := pjma.NewProject("", viper.GetString("context.project"))
			cobra.CheckErr(err)
			viper.Set("context.project", pj.Code())
			evars.AddReferProject(*pj)
		}

		// If password isn't specified, ask user it with blind
		if viper.GetString("context.user") != "" && viper.GetString("context.password") == "" {
			cmd.Print("Enter password: ")
			p, err := term.ReadPassword(int(syscall.Stdin))
			cobra.CheckErr(err)
			viper.Set("context.password", string(p))
			cmd.Println()
		}

		lnchr := pjma.NewLauncher(viper.GetString("projects_dir"), viper.GetString(batkey), args[0]).
			SetModule(viper.GetString("context.module")).
			SetTty(viper.GetBool("context.tty")).
			SetProject(viper.GetString("context.project")).
			SetUser(viper.GetString("context.user"), viper.GetString("context.password")).
			SetMdb(viper.GetString("context.mdb")).
			SetMacro(viper.GetString("context.macro"))

		cobra.CheckErr(evars.Save())
		cmd.Println("Running app " + args[0] + "...")
		cmd.Println("  Module:", viper.GetString("context.module"))
		cmd.Println("  TTY Mode:", viper.GetBool("context.tty"))
		cmd.Println("  Project:", viper.GetString("context.project"))
		cmd.Println("  User:", viper.GetString("context.user"))
		cmd.Println("  MDB:", viper.GetString("context.mdb"))
		cmd.Println("  Macro:", viper.GetString("context.macro"))
		cobra.CheckErr(lnchr.Run())
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return pjma.SortStringKeys(viper.GetStringMapString("apps")), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	openCmd.Flags().StringP("module", "e", "", "module to enter")
	openCmd.Flags().BoolP("tty", "t", false, "launch with no gui mode")
	openCmd.Flags().StringP("project", "P", "", "project code or dir")
	openCmd.Flags().StringP("user", "u", "", "username")
	openCmd.Flags().StringP("password", "p", "", "password")
	openCmd.Flags().StringP("mdb", "M", "", "MDB name")
	openCmd.Flags().StringP("macro", "m", "", "macro file")
}
