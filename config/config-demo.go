package config

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
)

func GetConfigClient() (client config_client.IConfigClient) {
	// 创建服务器配置
	serverConfig := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}

	//创建客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           1000,
		NotLoadCacheAtStart: true,
		LogDir:              "/Users/lixiaoshuang/nacos/log",
		CacheDir:            "/Users/lixiaoshuang/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	// 创建客户端
	configClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})
	return configClient
}
