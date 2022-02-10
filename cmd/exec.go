/*
Copyright © 2021 K.Awata <awata_kihachi@outlook.jp>

*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:     "exec SCRIPT",
	Aliases: []string{"x"},
	Short:   "Run a script defined by config file",
	Long:    "Run a script defined by config file",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		scrname := "scripts." + args[0]
		if !viper.IsSet(scrname) {
			fmt.Fprintln(os.Stderr, "script name not found in config file")
			return
		}

		var buf bytes.Buffer
		var quote rune
		words := []string{}
		esc := false
		for _, r := range viper.GetString(scrname) {
			if !esc && r == '\\' {
				esc = true
				continue
			}
			if esc {
				if r != '"' && r != '\'' && r != '\\' {
					buf.WriteRune('\\')
				}
				buf.WriteRune(r)
				esc = false
				continue
			}
			if r == quote {
				quote = 0
				continue
			}
			if quote == 0 && (r == '"' || r == '\'') {
				quote = r
				continue
			}
			if quote == 0 && r == ' ' {
				if buf.Len() > 0 {
					words = append(words, buf.String())
					buf.Reset()
				}
				continue
			}
			buf.WriteRune(r)
		}
		words = append(words, buf.String())

		words = append(words, args[1:]...)
		fmt.Println("> " + strings.Join(words, " "))
		out, err := exec.Command(words[0], words[1:]...).Output()
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
