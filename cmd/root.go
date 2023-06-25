package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "networkexplorer",
	Short: "NetworkExplorer helps evaluate various network parameters",
	Long:  `A simple CLI tool for evaluating various network parameters. Use NetworkExplorer to stay current about different network properties of any website with just the URL`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// Do Stuff Here
	// },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
