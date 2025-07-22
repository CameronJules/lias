/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/CameronJules/lias/alias"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a specifed alias from a project",
	Long: `Specify the project and alias from the project you want to run
	
	Example: lias run projectName functionName`,
	Run: runFunction,
}

// Execute the alias specified by the users input project and function alias
// Read in all of the projects and search for the project
func runFunction(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatalf(`
	Error: incorrect number of arguments for run command.

	Details:
		Command: "lias run" 
		Reason: expected 2 args, instead got: %v
	Tip:
		Format should be "lias run projectName FunctionName"`, len(args))
	}
	project_name := args[0]
	function_name := args[1]

	projects, err := alias.ReadProjects(viper.GetString("datafile"))
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var userCommand string
	projectFound := false
	for _, proj := range projects {
		if proj.Name == project_name {
			projectFound = true
			userCommand = proj.Functions[function_name]
		}
	}

	if !projectFound {
		log.Fatalf("Error: project %v not found.", project_name)
	}
	if userCommand == "" {
		log.Fatalf("Error: command %v found.", function_name)
	}

	command := exec.Command("sh", "-c", userCommand)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		log.Printf("error: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
