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

// listFunctionsCmd represents the listFunctions command
var listFunctionsCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `Usage: `,
	Run:   listFunctions,
}

func listFunctions(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatalf(`
	Error: Incorrect number of arguments for list project command
	
	Details: Expected 1, instead got: %v`, len(args))
	}

	project_name := args[0]
	projects, err := alias.ReadProjects(viper.GetString("datafile"))

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	found := false
	var project alias.Project
	for _, proj := range projects {
		if proj.Name == project_name {
			found = true
			project = proj
		}
	}

	if !found {
		log.Fatalf("Error: project %v not found", project_name)
	}

	for key, function := range project.Functions {
		fmt.Println("alias:", key, "  Command:", function)
	}

}

func init() {
	functionCmd.AddCommand(listFunctionsCmd)
	functionCmd.AddCommand()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listFunctionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listFunctionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
