/*
Copyright © 2021 K.Awata <awata_kihachi@outlook.jp>

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
		appname := "apps." + args[0]
		if !viper.IsSet(appname) {
			fmt.Fprintln(os.Stderr, "app name not found in config file")
			return
		}

		viper.Set("appname", appname)
		viper.BindPFlag("context.module", cmd.Flags().Lookup("module"))
		viper.BindPFlag("context.tty", cmd.Flags().Lookup("tty"))
		viper.BindPFlag("context.project", cmd.Flags().Lookup("project"))
		viper.BindPFlag("context.user", cmd.Flags().Lookup("user"))
		viper.BindPFlag("context.mdb", cmd.Flags().Lookup("mdb"))
		viper.BindPFlag("context.macro", cmd.Flags().Lookup("macro"))

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
