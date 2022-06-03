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
