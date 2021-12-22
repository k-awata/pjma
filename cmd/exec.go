/*
Copyright © 2021 K.Awata <awata_kihachi@outlook.jp>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec SCRIPT",
	Short: "Run a script defined by config file",
	Long:  "Run a script defined by config file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		scrname := "scripts." + args[0]
		if !viper.IsSet(scrname) {
			fmt.Fprintln(os.Stderr, "script name not found in config file")
			return
		}
		scr := append(strings.Fields(viper.GetString(scrname)), args[1:]...)
		fmt.Println("> " + strings.Join(scr, " "))
		out, err := exec.Command(scr[0], scr[1:]...).Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if len(out) > 0 {
			fmt.Print(string(out))
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
