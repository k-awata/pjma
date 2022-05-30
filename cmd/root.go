/*
Copyright © 2021 K.Awata <awata_kihachi@outlook.jp>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultconf = "pjmaconf.yaml"

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "pjma",
	Short:   "Project manager for Aveva E3D Design and Administration",
	Long:    "Project manager for Aveva E3D Design and Administration",
	Version: "0.5.0",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./"+defaultconf+")")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME")
		viper.SetConfigType("yaml")
		viper.SetConfigName("pjmaconf")
	}
	viper.ReadInConfig()
}
