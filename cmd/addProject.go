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

// addProjectCmd represents the addProject command
var addProjectCmd = &cobra.Command{
	Use:   "add",
	Short: "add projects",
	Long: `Add a new project to store aliases
	
	Example: lias project add projectName"`,
	Run: addProject,
}

func addProject(cmd *cobra.Command, args []string) {
	projects, err := alias.ReadProjects(viper.GetString("datafile"))
	if err != nil {
		if err.Error() != "no such file or directory" {
			log.Printf("%v", err)
		} else {
			fmt.Printf("Attempting save to: %v", viper.GetString("datafile"))
		}

	}
	name := args[0]

	project := alias.Project{Name: name, Functions: make(map[string]string)}
	projects = append(projects, project)

	if err := alias.SaveProjects(viper.GetString("datafile"), projects); err != nil {
		fmt.Errorf("%v", err)
	}
}

func init() {
	projectCmd.AddCommand(addProjectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addProjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addProjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
