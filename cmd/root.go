package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "api-crawler",
		Short: "crawls sites and saves items in db",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(addPlugin)
	rootCmd.AddCommand(crawl)
}
