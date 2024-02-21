package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cmd represents the command-line options
type cmd struct {
	url         string // URL to process
	concurrency int    // Number of concurrent requests
	rps         int    // Requests per second
}

func main() {
	// Define myCmd to hold command-line flags
	myCmd := &cmd{}
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
        Use:   "forge",
        Short: "Forge - Load Testing Utility",
        Long: `Forge is a load testing utility that can be used to test the performance of web applications.`,
        
		Run: func(cmd *cobra.Command, args []string) {
			// Execute the load test with provided parameters
			LoadTest(myCmd.url, myCmd.concurrency, myCmd.rps)
		},
	}

	// Define command-line flags
	rootCmd.Flags().StringVarP(&myCmd.url, "url", "u", "", "URL to process")
	rootCmd.Flags().IntVarP(&myCmd.concurrency, "concurrency", "n", 1, "Number of concurrent requests")
	rootCmd.Flags().IntVarP(&myCmd.rps, "rps", "r", 1, "Requests per second")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		// Print any errors encountered during command execution
		fmt.Println(err)
	}
}
