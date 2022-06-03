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
