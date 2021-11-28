package naming

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func GetNamingClient() naming_client.INamingClient {
	// 创建服务器配置
	serverConfig := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}

	//创建客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         "",
		TimeoutMs:           1000,
		NotLoadCacheAtStart: true,
		LogDir:              "/Users/lixiaoshuang/nacos/log",
		CacheDir:            "/Users/lixiaoshuang/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	// Create naming client for service discovery
	namingClient, _ := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})
	return namingClient
}

// Register 注册一个服务
func Register(ip string, port uint64) string {
	namingClient := GetNamingClient()
	result, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: "nacos-demo",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "cluster-a", // default value is DEFAULT
		GroupName:   "group-a",   // default value is DEFAULT_GROUP
	})
	if err != nil {
		fmt.Println(err)
		return "注册失败"
	}
	if result {
		return "注册成功"
	}
	return "注册失败"
}

// GetInstance 获取注册服务
func GetInstance(serviceName string, cluster string, group string) model.Service {
	namingClient := GetNamingClient()
	instance, _ := namingClient.GetService(vo.GetServiceParam{
		Clusters: []string{cluster}, ServiceName: serviceName,
		GroupName: group,
	})
	return instance
}
