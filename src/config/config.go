package config

import (
	"errors"
	"fmt"
	"os"
)

type configJson struct {
	RobotBaseUrl  string
	KubesphereUrl string
	DebugMode     bool
}

var Config = configJson{
	RobotBaseUrl: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send",
}

func LoadConfig() (err error) {
	kubesphereUrl := os.Getenv("KUBESPHERE_URL")
	if kubesphereUrl == "" {
		err = errors.New("环境变量中未获取到变量'KUBESPHERE_URL'")
	} else {
		Config.KubesphereUrl = kubesphereUrl
	}
	debugMode := os.Getenv("DEBUG")
	if debugMode != "" && debugMode != "false" && debugMode != "0" {
		Config.DebugMode = true
	}
	fmt.Printf("%+v\n", Config)
	return err
}

func init() {
	LoadConfig()
}
