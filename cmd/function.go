/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// functionCmd represents the function command
var functionCmd = &cobra.Command{
	Use:   "function",
	Short: "Function commands",
	Long:  `Usage: `,
}

func init() {
	rootCmd.AddCommand(functionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// functionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// functionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
