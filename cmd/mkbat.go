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

var mkbatCmd = &cobra.Command{
	Use:     "mkbat BATFILE APPNAME",
	Short:   "Make a batch file to run specified app",
	Long:    "Make a batch file to run specified app",
	Example: `pjma mkbat launch.bat e3d31`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		appname := "apps." + args[1]
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
		viper.BindPFlag("absbat", cmd.Flags().Lookup("abs"))

		if err := pjma.MakeLaunch(args[0]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println("batch file created at " + args[0])
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