package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	appHost     string
	appPort     string
	appConfPath string
)

var appCmd = &cobra.Command{
	Use:                        "app",
	Short:                      "启动门户网站服务",
	Long:                       `启动门户网站服务`,
	SuggestionsMinimumDistance: 10,
	SuggestFor:                 []string{"web"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("门户网站服务已启动...")
		//fmt.Println("启动详情", "host:", appHost, "port:", appPort, "confPath:", appConfPath)
	},
}

func init() {
	appCmd.Flags().StringVarP(&appHost, "host", "t", "0.0.0.0", "指定主机地址")
	appCmd.Flags().StringVarP(&appPort, "port", "p", "8080", "指定端口")
	appCmd.Flags().StringVarP(&appConfPath, "conf", "c", "./config/config.yml", "指定配置文件路径")
}
