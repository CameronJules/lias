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

// deleteFunctionCmd represents the deleteFunction command
var deleteFunctionCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete function from a project",
	Long: `Specify the project and delete the function if it exists
	
	lias delete function projectName functionName`,
	Run: deleteFunction,
}

func deleteFunction(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatalf("delete function expects 2 args but received %v", len(args))
	}

	project_name := args[0]
	function_name := args[1]

	projects, err := alias.ReadProjects(viper.GetString("datafile"))

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var updatedProjects []alias.Project
	for _, proj := range projects {
		if proj.Name == project_name {
			delete(proj.Functions, function_name)
		}
		updatedProjects = append(updatedProjects, proj)
	}

	if err := alias.SaveProjects(viper.GetString("datafile"), updatedProjects); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func init() {
	functionCmd.AddCommand(deleteFunctionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteFunctionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteFunctionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
