package cmd

import (
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:                        "start",
	Short:                      "项目启动",
	Long:                       `项目启动命令`,
	SuggestionsMinimumDistance: 10,
	SuggestFor:                 []string{"run"},
}

func init() {
	startCmd.AddCommand(amdinCmd)
	startCmd.AddCommand(appCmd)
}
