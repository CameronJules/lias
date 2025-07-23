/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Set active project",
	Long:  `Set a project as active enabling function and run commands to operate.`,
	Run:   useProject,
}

func useProject(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatalf("Error: use command expects 1 argument")
	}

	project_name := args[0]
	viper.Set("active_project", project_name)
	viper.WriteConfig()
	fmt.Println("Active project set to:", project_name)

}

func init() {
	rootCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
