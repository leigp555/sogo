package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	appHost     string
	appPort     string
	appConfPath string
)
var filesNames = []string{"config", "settings"}

var appCmd = &cobra.Command{
	Use:                        "app",
	Short:                      "启动门户网站服务",
	Long:                       `启动门户网站服务`,
	SuggestionsMinimumDistance: 10,
	SuggestFor:                 []string{"web"},
	PreRun: func(cmd *cobra.Command, args []string) {
		readConf()
		//TODO:初始化
		//TODO:启动服务
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("门户网站服务已启动...")
	},
}

func init() {
	appCmd.Flags().StringVarP(&appHost, "host", "t", "0.0.0.0", "指定主机地址")
	appCmd.Flags().StringVarP(&appPort, "port", "p", "8080", "指定端口")
	appCmd.Flags().StringVarP(&appConfPath, "conf", "c", "./config/config.yml", "指定配置文件路径")
}

func readConf() {
	//加载配置文件
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	for i := 0; i < len(filesNames); i++ {
		v.SetConfigName(filesNames[i])
		if i == 0 {
			if err := v.ReadInConfig(); err != nil {
				fmt.Println(fmt.Errorf("没有找到配置文件: %s", err.Error()))
				os.Exit(1)
			}
		} else {
			if err := v.MergeInConfig(); err != nil {
				fmt.Println(fmt.Errorf("没有找到配置文件: %s", err.Error()))
				os.Exit(1)
			}
		}
	}

	//读取配置文件
	fmt.Println(v.GetString("HttpServer.app.Port"))
	fmt.Println(v.GetString("mysql.addr"))
}
