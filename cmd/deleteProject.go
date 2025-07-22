/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/CameronJules/lias/alias"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteProjectCmd represents the deleteProject command
var deleteProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: deleteProject,
}

func deleteProject(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatalf("Delete project takes 1 cinput but got %v", len(args))
	}

	project_to_delete := args[0]
	projects, err := alias.ReadProjects(viper.GetString("datafile"))

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var updatedProjects []alias.Project
	for _, proj := range projects {
		if proj.Name == project_to_delete {
			continue
		} else {
			updatedProjects = append(updatedProjects, proj)
		}
	}
	if err := alias.SaveProjects(viper.GetString("datafile"), updatedProjects); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func init() {
	deleteCmd.AddCommand(deleteProjectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteProjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteProjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
