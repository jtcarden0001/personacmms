/*
Copyright © 2024 John Carden john.carden.02@gmail.com
*/
package root

import (
	"os"

	"github.com/jtcarden0001/personacmms/cli/cmd/create"
	"github.com/jtcarden0001/personacmms/cli/cmd/delete"
	"github.com/jtcarden0001/personacmms/cli/cmd/get"
	"github.com/jtcarden0001/personacmms/cli/cmd/list"
	"github.com/jtcarden0001/personacmms/cli/cmd/update"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmmsctl",
	Short: "A brief description of your application",
	Long:  `cmmsctl is a cli tool for interacting with the personacmms application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(delete.DeleteCmd)
	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(update.UpdateCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmmsctl/config)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cmms/config" (without extension).
		viper.AddConfigPath(home + "/.cmms")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.SetDefault("server.ip", "localhost")
	viper.SetDefault("server.port", "8080")

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore
		} else {
			// Config file was found but another error was produced, ignore
		}
	}

	// Config file found and successfully parsed
}
