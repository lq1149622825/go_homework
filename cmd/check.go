package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"go_homework/utils"
	"os"

	"github.com/spf13/cobra"
)

var check = false
var reload = false

func reloadMyProxy(myProxy **viper.Viper) {
	*myProxy = utils.GetConfigInfo()
}

func CheckConfigFormat() bool {
	if utils.GetConfigInfo().GetString("proxy_server.location.proxy_pass") == "" || utils.GetConfigInfo().GetString("proxy_server.listen") == "" {
		return false
	} else {
		return true
	}
}

func checkConfig(cmd *cobra.Command, args []string) {
	if !CheckConfigFormat() {
		fmt.Println("config info error!")
		os.Exit(1)
	}
	if reload {
		reloadMyProxy(&utils.MyProxyInfo)
		fmt.Println("reload success")
	}
	if check {
		fmt.Println("check config success!")
	}
}

var checkCmd = &cobra.Command{
	Use: "check",
	Run: checkConfig,
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
