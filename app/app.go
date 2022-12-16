package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_homework/utils"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var proxy = utils.MyProxy

func Proxy(c *gin.Context) {
	remote, err := url.Parse(proxy.GetString("proxy_server.location.proxy_pass"))
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
	server.Any("", Proxy)
	err := server.Run(proxy.GetString("proxy_server.listen"))
	if err != nil {
		fmt.Println("server failure")
	}
}
