/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/CameronJules/lias/alias"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listProjectsCmd represents the listProjects command
var listProjectsCmd = &cobra.Command{
	Use:   "list",
	Short: "list projects",
	Long:  `List projects available`,
	Run:   listProjects,
}

func listProjects(cmd *cobra.Command, args []string) {
	projects, err := alias.ReadProjects(viper.GetString("datafile"))

	if err != nil {
		log.Printf("%v", err)
	}
	for i, project := range projects {
		fmt.Println("(", i, ") ", project.Name)
	}
}

func init() {
	projectCmd.AddCommand(listProjectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listProjectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listProjectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
