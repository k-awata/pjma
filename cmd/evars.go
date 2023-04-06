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
	"github.com/k-awata/pjma/pjma"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// evarsCmd represents the evars command
var evarsCmd = &cobra.Command{
	Use:   "evars",
	Short: "Update custom_evars.bat in projects_dir",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		evars := pjma.NewEvars(viper.GetString("projects_dir"))
		cobra.CheckErr(evars.AddReferProjectDirs(viper.GetStringSlice("refer_pj")))
		evars.AddJoinEnv(viper.GetStringMapStringSlice("join_env"))
		evars.AddAfterCmd(viper.GetString("after_cmd"))
		cobra.CheckErr(evars.Save(viper.GetString("encoding")))
		cmd.Println("pjma updated custom_evars.bat in projects_dir")
	},
}

func init() {
	rootCmd.AddCommand(evarsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// evarsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// evarsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
