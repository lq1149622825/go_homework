package cmd

import (
	"github.com/fatih/color"
	"go_homework/utils"
	"os"

	"github.com/spf13/cobra"
)

var check = false
var reload = false

var MyProxyInfo = utils.MyProxy

func CheckConfigFormat() bool {
	return (MyProxyInfo.GetString("proxy_server.location.proxy_pass") == "") && (MyProxyInfo.GetString("proxy_server.listen") != "")

}

func checkAndReloadInitConfig(cmd *cobra.Command, args []string) {
	colorPrint := color.New()
	if !CheckConfigFormat() {
		colorPrint.Add(color.FgRed)
		colorPrint.Println("config info error!")
		os.Exit(1)
	}
	if reload {
		//TODO 进程通信 告知app服务 需要ReadInConfig

		colorPrint.Add(color.FgGreen)
		colorPrint.Println("reload success!")
	}
	if check {
		colorPrint.Add(color.FgGreen)
		colorPrint.Println("check config success!")
	}
}

var checkCmd = &cobra.Command{
	Use: "check",
	Run: checkAndReloadInitConfig,
}

func init() {
	checkCmd.PersistentFlags().BoolVarP(&check, "check", "c", false, "check")
	checkCmd.PersistentFlags().BoolVarP(&reload, "reload", "r", false, "reload")
}

func Execute() {
	err := checkCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
