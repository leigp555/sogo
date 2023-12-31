package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:                        "sogo",
	Short:                      "sogo web脚手架",
	Long:                       `sogo 易扩展,功能强大的golang web脚手架`,
	Version:                    "0.0.1",
	SuggestionsMinimumDistance: 1,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(startCmd)
}
