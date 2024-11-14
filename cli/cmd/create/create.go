/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a resource",
	Run: func(cmd *cobra.Command, args []string) {
		// Print the help message if no subcommands are provided
		if err := cmd.Help(); err != nil {
			fmt.Println("Error showing help:", err)
		}
	},
}

func init() {
	// sub commands
	CreateCmd.AddCommand(createCategoryCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
