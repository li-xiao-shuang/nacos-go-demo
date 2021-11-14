package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"nacos-go-demo/config"
)

func main() {

	configClient := config.GetConfigClient()
	// 获取配置
	content, _ := configClient.GetConfig(vo.ConfigParam{DataId: "test-1", Group: "DEFAULT_GROUP"})
	fmt.Println("获取到的配置: " + content)


}
