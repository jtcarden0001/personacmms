/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package category

import (
	"fmt"

	"github.com/jtcarden0001/personacmms/cmmsctl/cmd"
	"github.com/spf13/cobra"
)

// deleteCategoryCmd represents the deleteCategory command
var deleteCategoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteCategory called")
	},
}

func init() {
	cmd.DeleteCmd.AddCommand(deleteCategoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCategoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCategoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
