package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	adminHost     string
	adminPort     string
	adminConfPath string
)

// amdinCmd represents the amdin command
var amdinCmd = &cobra.Command{
	Use:                        "admin",
	Short:                      "启动后台管理系统服务",
	Long:                       `启动后台管理系统服务`,
	SuggestionsMinimumDistance: 10,
	SuggestFor:                 []string{"server"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("后台管理系统服务启动...")
		//fmt.Println("host:", adminHost, "port:", adminPort, "confPath:", adminConfPath)
	},
}

func init() {
	amdinCmd.Flags().StringVarP(&adminHost, "host", "t", "0.0.0.0", "指定主机地址")
	amdinCmd.Flags().StringVarP(&adminPort, "port", "p", "8080", "指定端口")
	amdinCmd.Flags().StringVarP(&adminConfPath, "conf", "c", "./config/config.yml", "指定配置文件路径")
}
