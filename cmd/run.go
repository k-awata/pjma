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
	"os"
	"os/exec"

	"github.com/k-awata/pjma/pjma"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run script_name [args]...",
	Short: "Run a script defined by pjma.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// List available scripts
			fmt.Println("Available Scripts:")
			s := viper.GetStringMapString("scripts")
			for _, v := range pjma.SortStringKeys(s) {
				fmt.Println("  " + v)
				fmt.Println("    " + s[v])
			}
			return
		}

		// Read script
		scrkey := "scripts." + args[0]
		if !viper.IsSet(scrkey) {
			cobra.CheckErr(fmt.Errorf("unknown command %q for %q", args[0], cmd.CommandPath()))
		}
		scrval := append(pjma.ParseCommand(viper.GetString(scrkey)), args[1:]...)

		// Run script
		e := exec.Command(scrval[0], scrval[1:]...)
		e.Stdin = os.Stdin
		e.Stdout = os.Stdout
		e.Stderr = os.Stderr
		if err := e.Run(); err != nil {
			cmd.Println(err)
		}
		os.Exit(e.ProcessState.ExitCode())
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return pjma.SortStringKeys(viper.GetStringMapString("scripts")), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
