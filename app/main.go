package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var myProxy = getConfigInfo()

func getConfigInfo() *viper.Viper {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config := viper.New()
	config.AddConfigPath(path)     //设置读取的文件路径
	config.SetConfigName("config") //设置读取的文件名
	config.SetConfigType("yaml")   //设置文件的类型

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	return config
}

func Proxy(c *gin.Context) {
	remote, err := url.Parse(myProxy.GetString("proxy_server.location.proxy_pass"))
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	server := gin.Default()
	server.GET("", Proxy)
	err := server.Run(myProxy.GetString("proxy_server.listen"))
	if err != nil {
		fmt.Println("server failure")
	}
}
