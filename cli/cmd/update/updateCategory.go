/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCategoryCmd represents the updateCategory command
var updateCategoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateCategory called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCategoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCategoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
