/*
Copyright © 2021 K.Awata <awata_kihachi@outlook.jp>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/k-awata/pjma/internal/pjma"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:     "exec SCRIPT",
	Aliases: []string{"x"},
	Short:   "Run a script defined by config file",
	Long:    "Run a script defined by config file",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Available Scripts:")
			scr := viper.GetStringMapString("scripts")
			for _, v := range pjma.StringMapKeysToSlice(viper.GetStringMap("scripts")) {
				fmt.Println("  " + v)
				fmt.Println("    " + scr[v])
			}
			return
		}
		scrname := "scripts." + args[0]
		if !viper.IsSet(scrname) {
			fmt.Fprintln(os.Stderr, "script name not found in config file")
			return
		}
		scr := append(pjma.ParseScript(viper.GetString(scrname)), args[1:]...)
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
