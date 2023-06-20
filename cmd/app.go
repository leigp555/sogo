package cmd

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"sogo/app/global/variable"
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
		//读取配置文件
		readConf()
		//初始化依赖
		//初始化路由
		//启动项目
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
		variable.Config = v
	}
	// 监控配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		//do something
		fmt.Println("配置文件已更新")
	})
}
