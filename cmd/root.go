/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var dataFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lias",
	Short: "lias alias manager",
	Long: `
Lias allows you to organize your aliases into projects.
It enables reuse of alias names and simple management for addition 
and deletion of aliases and projects.

(Over-engineered aliases to learn Golang)
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)

	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set datafile using --datafile")
	}

	rootCmd.PersistentFlags().StringVar(
		&dataFile,
		"datafile",
		home+string(os.PathSeparator)+".lias.json",
		"data file to store projects")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config/.lias.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.BindPFlag("datafile", rootCmd.PersistentFlags().Lookup("datafile"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".lias")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("lias")

	viper.SetDefault("active_project", "")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config:", err)
	}

}
