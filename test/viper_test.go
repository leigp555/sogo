package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"testing"
)

func TestViper(t *testing.T) {
	server()
}

func server() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		s := viperRead()
		c.JSON(http.StatusOK, gin.H{"msg": "hello", "port": s})
	})
	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func viperRead() (s string) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("../config")
	if err := v.ReadInConfig(); err != nil {
		fmt.Println(fmt.Errorf("没有找到配置文件: %s", err.Error()))
		os.Exit(1)
	}
	return v.GetString("HttpServer.app.Port")
}
