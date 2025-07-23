/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// activeCmd represents the active command
var activeCmd = &cobra.Command{
	Use:   "active",
	Short: "Show current active project",
	Long:  `Usage: `,
	Run: func(cmd *cobra.Command, args []string) {
		project := viper.GetString("active_project")
		if project == "" {
			fmt.Println("No active project set.")
		} else {
			fmt.Printf("Active project: %s", project)
		}
	},
}

func init() {
	rootCmd.AddCommand(activeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// activeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// activeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
