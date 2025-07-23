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

// addFunctionCmd represents the addFunction command
var addFunctionCmd = &cobra.Command{
	Use:   "add",
	Short: "Add function to project",
	Long: `
	Add function to a specified project
	
	Example: lias add function proj_name alias function`,
	Run: addFunction,
}

func addFunction(cmd *cobra.Command, args []string) {
	active_project := viper.GetString("active_project")

	if active_project == "" {
		log.Fatal("Active project not set. run the command 'lias use [project] to set a project")
	}

	if len(args) != 2 {
		log.Fatalln("lias expected 2 arguments but got:", len(args))
	}
	function_alias := args[0]
	function_string := args[1]

	projects, err := alias.ReadProjects(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	for _, project := range projects {
		if project.Name == active_project {
			project.Functions[function_alias] = function_string
		}
	}

	if err := alias.SaveProjects(viper.GetString("datafile"), projects); err != nil {
		log.Printf("%v", err)
	}
}

func init() {
	functionCmd.AddCommand(addFunctionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addFunctionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addFunctionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
