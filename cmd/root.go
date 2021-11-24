package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ant1k9/api-crawler/config"
)

var (
	rootCmd = &cobra.Command{
		Use:   "api-crawler",
		Short: "crawls sites and saves items in db",
	}

	cfg config.Config
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(addPlugin)
	rootCmd.AddCommand(crawl)
	cfg = config.Init()
}
