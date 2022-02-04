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

// mkevarsCmd represents the mkevars command
var mkevarsCmd = &cobra.Command{
	Use:   "mkevars",
	Short: "Make custom_evars.bat in projects_dir",
	Long:  "Make custom_evars.bat in projects_dir",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if !viper.IsSet("projects_dir") {
			fmt.Fprintln(os.Stderr, "projects_dir is undefined")
			return
		}
		if err := pjma.MakeEvars(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println(viper.GetString("projects_dir") + `\custom_evars.bat has been updated`)
	},
}

func init() {
	rootCmd.AddCommand(mkevarsCmd)
}
