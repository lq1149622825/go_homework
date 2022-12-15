package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_homework/utils"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Proxy(c *gin.Context) {
	remote, err := url.Parse(utils.MyProxyInfo.GetString("proxy_server.location.proxy_pass"))
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
	fmt.Println(utils.MyProxyInfo.GetString("proxy_server.location.proxy_pass"))
	server := gin.Default()
	server.Any("", Proxy)
	err := server.Run(utils.MyProxyInfo.GetString("proxy_server.listen"))
	if err != nil {
		fmt.Println("server failure")
	}
}
